package etcdplug

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/boxcolli/pepperlink/plugs"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
)

var (
	ErrInvalidEtcdOption = errors.New("invalid etcd option")
	ErrEtcdDeadlineExceeded = errors.New("deadline exceeded while creating a etcd server")
)

type EtcdOption struct {
	// Cluster id to the new cluster.
	// Set with empty string if this etcd server is to be merged with existing cluster.
	Cid string

	// Data directory of etcd server.
	// Set with empty string for default directory (/var/lib/etcd)
	Dir string

	// Name of the new etcd server.
	Name string

	// Set with empty string if this etcd server is initializing a new cluster.
	PeerName string
	PeerUrl string

	// Policy for client request timeout.
	DialTimeout time.Duration

	// Only supports debug, info, warn, error, panic, or fatal. Default 'info'.
	LogLevel string
}

func initEtcd(ctx context.Context, opt EtcdOption) (*embed.Etcd, *clientv3.Client, error) {
	// Validate option
	{
		if opt.Cid == "" && (opt.PeerName == "" || opt.PeerUrl == "") {
			return nil, nil, ErrInvalidEtcdOption
		}
	}

	// Create etcd server
	var etcd *embed.Etcd
	{
		var err error

		// Configurations
		cfg := embed.NewConfig()		
		cfg.Dir = opt.Dir
		cfg.Name = opt.Name
		cfg.ListenPeerUrls = []url.URL{{Scheme: "http", Host: "0.0.0.0:0"}}
		cfg.ListenClientUrls = []url.URL{{Scheme: "http", Host: "0.0.0.0:0"}}
		if opt.PeerName != "" {
			cfg.InitialCluster = fmt.Sprintf("%s=%s,%s=%s", opt.Name, "http://localhost:0", opt.PeerName, opt.PeerUrl)
			cfg.AdvertisePeerUrls = []url.URL{{Scheme: "http", Host: opt.PeerUrl}}
		} else {
			cfg.InitialCluster = fmt.Sprintf("%s=%s", opt.Name, "http://localhost:0")
			cfg.AdvertisePeerUrls = []url.URL{{Scheme: "http", Host: "localhost:0"}}
		}
		cfg.LogLevel = opt.LogLevel
		
		// Start etcd
		etcd, err = embed.StartEtcd(cfg)
		if err != nil {
			return nil, nil, err
		}

		// Wait for the result
		select {
		case <- etcd.Server.ReadyNotify():
			break
		case <- ctx.Done():
			etcd.Server.Stop()
			return nil, nil, ErrEtcdDeadlineExceeded
		}
	}
	
	// Create etcd client
	var client *clientv3.Client
	{
		var err error

		// Create new etcd client
		client, err = clientv3.New(clientv3.Config{
			Endpoints: etcd.Server.Cfg.ClientURLs.StringSlice(),
			DialTimeout: opt.DialTimeout,
		})
		if err != nil {
			return nil, nil, err
		}

		// Set cluster id if necessary.
		if opt.Cid != "" {
			kv := clientv3.NewKV(client)
			_, err = kv.Put(ctx, plugs.KeyCid, opt.Cid)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return etcd, client, nil
}
