package etcdplug

import (
	"sync"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
)

const (
	KeyPubCount  = "pubcount"
	KeyPubPrefix = "pub"
)

type etcdPlug struct {
	etcd    *embed.Etcd // for etcd.Clos()
	cli     *clientv3.Client

	pubs    map[string]bool // local cache of publisher set
	pubsMux sync.RWMutex
}

// WatchPub implements cluster.Cluster.
func (*etcdPlug) WatchPub() {
	panic("unimplemented")
}

// Destroy implements cluster.Cluster.
func (c *etcdPlug) Destroy() {
	c.etcd.Server.Stop()
}
