package etcdplug

import (
	"context"
	"log"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/boxcolli/pepperlink/plugs"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
	"google.golang.org/grpc/connectivity"
)

func TestEtcdPlug(t *testing.T) {
	var etcd, client = createEtcdServer()
	var members = []plugs.Member{}
	var prefix = "local"
	var cname = "c0"
	{
		if etcd == nil {
			t.Fatal("cannot create etcd embed server.")
			t.FailNow()
		}

		// Clear the etcd database
		_, err := client.Delete(context.Background(), "", clientv3.WithPrefix())
		assert.NoError(t, err)

		defer etcd.Server.Stop()
	}

	var plug plugs.Plug
	{
		plug = NewEtcdPlug(client, prefix)
	}

	// Read
	{
		_, err := client.Put(context.Background(), "local:c0:n0", "0:0")
		assert.NoError(t, err)

		members = append(members, plugs.Member{
			Cname: "c0",
			Name: "n0",
			Host: "0",
			Port: "0",
		})
		ms, err := plug.Read(context.Background(), cname)
		assert.NoError(t, err)
		t.Log("ms:", ms)
		assert.Equal(t, true, reflect.DeepEqual(members, ms))
	}

	// Watch & Stop
	{
		change := plug.Watch(context.Background(), cname, 100)
		cs := []plugs.Change{}
		go func() {
			for ch := range change {
				cs = append(cs, ch)
			}
		}()

		changes := []plugs.Change{
			{
				Method: plugs.MethodPut,
				Data: plugs.Member{
					Cname: "c0",
					Name: "n1",
					Host: "1",
					Port: "1",
				},
			},
			{
				Method: plugs.MethodDel,
				Data: plugs.Member{
					Cname: "c0",
					Name: "n1",
					Host: "",
					Port: "",
				},
			},
		}
		_, err := client.Put(context.Background(), "local:c0:n1", "1:1")
		assert.NoError(t, err)
		_, err = client.Delete(context.Background(), "local:c0:n1")
		assert.NoError(t, err)

		time.Sleep(time.Second * 1)	// It takes time to detect changes
		plug.Stop(cname)
		t.Log("cs:", cs)
		assert.Equal(t, true, reflect.DeepEqual(changes, cs))
	}

	// Advertise
	{
		change := plug.Watch(context.Background(), cname, 100)
		cs := []plugs.Change{}
		go func() {
			for ch := range change {
				cs = append(cs, ch)
			}
		}()

		m := plugs.Member{
			Name: "2",
			Host: "2",
			Port: "2",
		}
		err := plug.Advertise(context.Background(), cname, m)
		assert.NoError(t, err)

		plug.Stop(cname)
		assert.Zero(t, len(cs))
	}

	// Destroy
	{
		plug.Destroy()
		
		conn := client.ActiveConnection()
		assert.Equal(t, connectivity.Shutdown, conn.GetState())
	}
}

func createEtcdServer() (*embed.Etcd, *clientv3.Client) {
	var cfg = embed.NewConfig()
	{
		cfg.Name = "local"
		cfg.Dir = "./dir/"
		cfg.ListenPeerUrls = parseUrls([]string{"http://0.0.0.0:2380"})
		cfg.ListenClientUrls = parseUrls([]string{"http://0.0.0.0:2379"})
		cfg.AdvertisePeerUrls = parseUrls([]string{"http://localhost:2380"})
		cfg.AdvertiseClientUrls = parseUrls([]string{"http://localhost:2379"})
		cfg.InitialCluster = "local=http://localhost:2380"
		cfg.LogLevel = "error"
	}

	var etcd *embed.Etcd
	{
		var err error
		etcd, err = embed.StartEtcd(cfg)
		if err != nil {
			log.Fatal(err)
		}

		select {
		case <- etcd.Server.ReadyNotify():
			log.Printf("Server is ready!")
		case <- time.After(10 * time.Second):
			etcd.Server.Stop()
			log.Printf("Server took too long to start!")
			return nil, nil
		}
	}

	var client *clientv3.Client
	{
		var err error
		client, err = clientv3.New(clientv3.Config{
			DialTimeout: 10 * time.Second,
			Endpoints: []string{"localhost:2379"},
		})
		if err != nil {
			log.Fatal(err)
			etcd.Close()
			return nil, nil
		}
	}

	return etcd, client
}

func parseUrls(strs []string) []url.URL {
	urls := make([]url.URL, 0, len(strs))
	for _, str := range strs {
		u, err := url.Parse(str)
		if err != nil {
			log.Printf("Invalid url %s, error: %s", str, err.Error())
			continue
		}
		urls = append(urls, *u)
	}

	return urls
}