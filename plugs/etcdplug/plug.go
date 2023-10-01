package etcdplug

import (
	"context"
	"sync"

	"github.com/boxcolli/pepperlink/plugs"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
)

const (
	KeyPubCount  = "pubcnt"
	KeyPubPrefix = "pub"
)

type etcdPlug struct {
	etcd *embed.Etcd // for etcd.Clos()
	cli  *clientv3.Client

	pubs    map[string]bool // local cache of publisher set
	pubsMux sync.RWMutex
}

func NewEtcdPlug(ctx context.Context, opt EtcdOption) (plugs.Plug, error) {
	p := &etcdPlug{}
	{
		var err error
		p.etcd, p.cli, err = initEtcd(ctx, opt)
		if err != nil {
			return nil, err
		}

	}
	return p, nil
}

// GetDiscoveryAddr implements plugs.Plug.
func (*etcdPlug) GetDiscoveryAddr() {
	panic("unimplemented")
}

// WatchPub implements cluster.Cluster.
func (*etcdPlug) WatchPub() {
	panic("unimplemented")
}

// Destroy implements cluster.Cluster.
func (c *etcdPlug) Destroy() {
	c.etcd.Server.Stop()
}
