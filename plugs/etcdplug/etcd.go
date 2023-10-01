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
	cid string

	// Data directory of etcd server.
	// Set with empty string for default directory (/var/lib/etcd)
	dir string

	// Name of the new etcd server.
	name string

	// Set with empty string if this etcd server is initializing a new cluster.
	peerName string
	peerUrl string

	// Policy for client request timeout.
	dialTimeout time.Duration

	// Only supports debug, info, warn, error, panic, or fatal. Default 'info'.
	logLevel string
}

func initEtcd(ctx context.Context, opt EtcdOption) (*embed.Etcd, *clientv3.Client, error) {
	// Validate option
	{
		if opt.cid == "" && (opt.peerName == "" || opt.peerUrl == "") {
			return nil, nil, ErrInvalidEtcdOption
		}
	}

	// Create etcd server
	var etcd *embed.Etcd
	{
		var err error

		// Configurations
		cfg := embed.NewConfig()		
		cfg.Dir = opt.dir
		cfg.Name = opt.name
		cfg.ListenPeerUrls = []url.URL{{Scheme: "http", Host: "127.0.0.1:0"}}
		cfg.ListenClientUrls = []url.URL{{Scheme: "http", Host: "127.0.0.1:0"}}
		if opt.peerName != "" {
			cfg.InitialCluster = fmt.Sprintf("%s=%s,%s=%s")
			cfg.AdvertisePeerUrls = []url.URL{{Scheme: "http", Host: opt.peerUrl}}
		}
		cfg.LogLevel = opt.logLevel
		
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

		client, err = clientv3.New(clientv3.Config{
			Endpoints: etcd.Server.Cfg.ClientURLs.StringSlice(),
			DialTimeout: opt.dialTimeout,
		})
		if err != nil {
			return nil, nil, err
		}

		// Set cluster id if necessary.
		if opt.cid != "" {
			kv := clientv3.NewKV(client)
			_, err = kv.Put(ctx, plugs.KeyCid, opt.cid)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return etcd, client, nil
}
