package etcdplug

import (
	"context"
	"log"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
	"google.golang.org/grpc/connectivity"
)

func TestEtcdPlug(t *testing.T) {
	var etcd, client = createEtcdServer()
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

	var (
		prefix = "local"
		delim = "#"
		cname = "c0"
		plug plugs.Plug
		formatter plugs.Formatter
	)
	{
		formatter = plugs.NewBasicFormatter(prefix, delim)
		plug = NewEtcdPlug(client, formatter)
	}

	var members = []types.Member{}
	{
		members = append(members, types.Member{
			Cname: "c0",
			Name: "n0",
			Pro: types.ProtocolGrpc,
			Host: "0",
			Port: "0",
		})
	}
	{
		_, err := client.Put(
			context.Background(),
			formatter.PrintKey(&members[0]),
			formatter.PrintValue(&members[0]),
		)
		assert.NoError(t, err)
	}

	// Watch & Stop
	{
		// Watch
		ch, err := plug.Watch(context.Background(), cname, 100)
		assert.NoError(t, err)
		es := []plugs.Event{}
		go func() {
			for e := range ch {
				es = append(es, *e)
			}
		} ()

		// Expected
		m := types.Member{
			Cname: "c0",
			Name: "n1",
			Pro: types.ProtocolGrpc,
			Host: "1",
			Port: "1",
		}
		members = append(members, m)
		events := []plugs.Event{
			{
				Op: types.OperationAdd,
				Data: members[0],
			},
			{
				Op: types.OperationAdd,
				Data: m,
			},
			{
				Op: types.OperationDel,
				Data: &types.Member{
					Cname: m.Cname,
					Name: m.Name,
				},
			},
		}

		// Put member
		_, err = client.Put(
			context.Background(),
			formatter.PrintKey(&m),
			formatter.PrintValue(&m),
		)
		assert.NoError(t, err)

		// Delete member
		_, err = client.Delete(context.Background(), formatter.PrintKey(&m))
		assert.NoError(t, err)

		// Assert
		time.Sleep(time.Second * 2)	// It takes time to detect changes
		plug.Stop(cname)
		assert.Equal(t, true, reflect.DeepEqual(events, es))
		t.Log("watch actual events", es)
		t.Log("watch actual events len(es)", len(es))
	}

	// Clear the etcd database
	_, err := client.Delete(context.Background(), "", clientv3.WithPrefix())
	assert.NoError(t, err)

	// Me
	{
		// Watch
		es := []plugs.Event{}
		{
			ch, err := plug.Watch(context.Background(), cname, 100)
			assert.NoError(t, err)
			go func() {
				for e := range ch {
					es = append(es, *e)
				}
			}()
		}

		me := &types.Member{
			Cname: "c",
			Name: "2",
			Pro: types.ProtocolGrpc,
			Host: "2",
			Port: "2",
		}

		// Add me
		{
			err := plug.Me(context.Background(), types.OperationAdd, me)
			assert.NoError(t, err)
	
			// Assert
			plug.Stop(cname)
			time.Sleep(1 * time.Second)
			assert.Zero(t, len(es))
			t.Log("Add me watch events",es)
	
			res, err := client.Get(context.Background(), formatter.PrintKey(me))
			assert.NoError(t, err)
			for _, kv := range res.Kvs {
				m := new(types.Member)
				formatter.ScanKey(string(kv.Key), m)
				formatter.ScanValue(string(kv.Value), m)
				assert.Equal(t, true, reflect.DeepEqual(me, m))
			}
		}

		// Delete me
		{
			err := plug.Me(context.Background(), types.OperationDel, nil)
			assert.NoError(t, err)
			
			res, err := client.Get(context.Background(), formatter.PrintKey(me))
			assert.NoError(t, err)
			assert.Zero(t, len(res.Kvs))
		}
	}

	// Close
	{
		plug.Close()
		
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
