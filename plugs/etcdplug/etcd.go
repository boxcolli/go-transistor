package etcdplug

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"net"
// 	"net/url"
// 	"time"

// 	"github.com/boxcolli/pepperlink/plugs"
// 	"go.etcd.io/etcd/client/v3"
// 	"go.etcd.io/etcd/server/v3/embed"
// )

// var (
// 	ErrInvalidEtcdOption = errors.New("invalid etcd option")
// 	ErrEtcdDeadlineExceeded = errors.New("deadline exceeded while creating a etcd server")
// )

// type EtcdOption struct {
// 	// Cluster id to the new cluster.
// 	// Set with empty string if this etcd server is to be merged with existing cluster.
// 	Cid string

// 	// Data directory of etcd server.
// 	// Set with empty string for default directory (/var/lib/etcd)
// 	Dir string

// 	// Name of the new etcd server and discovery address.
// 	Name string
// 	Addr string

// 	// Set with empty string if this etcd server is initializing a new cluster.
// 	PeerName string
// 	PeerAddr string

// 	// Policy for client request timeout.
// 	DialTimeout time.Duration

// 	// Only supports debug, info, warn, error, panic, or fatal. Default 'info'.
// 	LogLevel string

// 	// new, existing
// 	ClusterState string
// }

// func initEtcd(ctx context.Context, client *clientv3.Client, opt EtcdOption) (error) {
// 	// Validate option
// 	{
// 		if opt.Cid == "" && (opt.PeerName == "" || opt.PeerAddr == "") {
// 			return nil, nil, ErrInvalidEtcdOption
// 		}
// 	}

// 	// Get URL
// 	var lpu *url.URL
// 	var lcu *url.URL
// 	var apu *url.URL
// 	var acu *url.URL
// 	{
// 		var err error

// 		// Generate URL for peer
// 		pp, err := getAvailablePort()
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		lpu, err = getURL("http", "0.0.0.0", pp)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		apu, err = getURL("http", "localhost", pp)
// 		if err != nil {
// 			return nil, nil, err
// 		}

// 		// Generate URL for client
// 		cp, err := getAvailablePort()
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		lcu, err = getURL("http", "0.0.0.0", cp)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		acu, err = getURL("http", "localhost", cp)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 	}
	

// 	// Create etcd server
// 	var etcd *embed.Etcd
// 	{
// 		var err error

// 		// Configurations
// 		cfg := embed.NewConfig()		
// 		cfg.Dir = opt.Dir
// 		cfg.Name = opt.Name
// 		cfg.ListenPeerUrls = []url.URL{*lpu}
// 		cfg.ListenClientUrls = []url.URL{*lcu}
// 		if opt.PeerName != "" {
// 			cfg.InitialCluster = fmt.Sprintf("%s=%s,%s=%s", opt.Name, "http://localhost:0", opt.PeerName, opt.PeerUrl)
// 			cfg.AdvertisePeerUrls = []url.URL{{Scheme: "http", Host: opt.PeerUrl}}
// 		} else {
// 			cfg.InitialCluster = fmt.Sprintf("%s=%s", opt.Name, "http://localhost:0")
// 			cfg.AdvertisePeerUrls = []url.URL{{Scheme: "http", Host: "localhost:0"}}
// 		}
// 		cfg.LogLevel = opt.LogLevel
// 		cfg.ClusterState = opt.ClusterState
		
// 		// Start etcd
// 		etcd, err = embed.StartEtcd(cfg)
// 		if err != nil {
// 			return nil, nil, err
// 		}

// 		// Wait for the result
// 		select {
// 		case <- etcd.Server.ReadyNotify():
// 			break
// 		case <- ctx.Done():
// 			etcd.Server.Stop()
// 			return nil, nil, ErrEtcdDeadlineExceeded
// 		}
// 	}
	
// 	// Create etcd client
// 	var client *clientv3.Client
// 	{
// 		var err error

// 		// Create new etcd client
// 		client, err = clientv3.New(clientv3.Config{
// 			Endpoints: etcd.Server.Cfg.ClientURLs.StringSlice(),
// 			DialTimeout: opt.DialTimeout,
// 		})
// 		if err != nil {
// 			return nil, nil, err
// 		}

// 		// Set cluster id if necessary.
// 		if opt.Cid != "" {
// 			kv := clientv3.NewKV(client)
// 			_, err = kv.Put(ctx, plugs.KeyCid, opt.Cid)
// 			if err != nil {
// 				return nil, nil, err
// 			}
// 		}
// 	}

// 	return etcd, client, nil
// }

// func getAvailablePort() (string, error) {
// 	// Reserve a port
// 	lis, err := net.Listen("tcp", ":0")
// 	if err != nil {
// 		return "", err
// 	}

// 	// Extract port from reserved address
// 	addr := lis.Addr().String()
// 	i := len(addr) - 1
// 	for ; i >= 0; i-- {
// 		if addr[i] == ':' {
// 			break
// 		}
// 	}

// 	return addr[i + 1:], nil
// }

// func getURL(scheme, host, port string) (*url.URL, error) {
// 	// Generate URL
// 	u, err := url.Parse(fmt.Sprintf("%s://%s:%s", scheme, host, port))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return u, nil
// }