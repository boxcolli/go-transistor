package etcdplug

import (
	"context"
	"testing"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
)

func TestInitEtcd(t *testing.T) {
	var e1 *embed.Etcd
	var c1 *clientv3.Client
	{
		var opt EtcdOption
		opt.Cid = "c1"
		opt.Dir = "dir/e1/"
		opt.Name = "c1-n1"
		opt.LogLevel = "error"

		var err error
		e1, c1, err = initEtcd(context.Background(), opt)
		if err != nil {
			t.Fatal(err)
		}
		defer e1.Close()
	}

	lpu := e1.Config().ListenPeerUrls
	t.Log("LPU:", lpu)

	{
		kv := clientv3.NewKV(c1)
		_, err := kv.Put(context.Background(), "k", "v")
		if err != nil {
			t.Fatal(err)
		}
	}
}