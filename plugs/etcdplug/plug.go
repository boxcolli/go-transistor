package etcdplug

import (
	"context"
	"fmt"
	"sync"

	"github.com/boxcolli/pepperlink/plugs"
	"go.etcd.io/etcd/client/v3"
)

const (
	fKeyspace = "%s:%s:"
	fKey = "%s:%s:%s"
	fValue = "%s:%s"
)

type etcdPlug struct {
	client *clientv3.Client

	// Special key prefix for key name
	prefix string

	// To prevent emitting myself in watch channel
	me   *plugs.Member
	memx sync.RWMutex

	w   map[string]chan plugs.Change	// Singleton watch channels
	stopw map[string]chan bool         	// Channels connected with watch goroutines
	wmx sync.Mutex
}

func NewEtcdPlug(client *clientv3.Client, prefix string) plugs.Plug {
	plug := &etcdPlug{
		client: client,
		prefix: prefix,
		me: nil,
		w: map[string]chan plugs.Change{},
		stopw: map[string]chan bool{},
	}
	return plug
}

// Advertise implements plugs.Plug.
func (p *etcdPlug) Advertise(ctx context.Context, cname string, me plugs.Member) error {
	// Write myself on local
	p.memx.Lock()
	p.me = &me
	p.memx.Unlock()

	// Advertise myself on the KV store
	_, err := p.client.Put(ctx, p.key(cname, me.Name), p.value(me.Host, me.Port))
	if err != nil {
		return err
	}

	return nil
}

// Read implements plugs.Plug.
func (p *etcdPlug) Read(ctx context.Context, cname string) ([]plugs.Member, error) {
	// Get keys with prefix, cname
	res, err := p.client.Get(ctx, p.keyspace(cname), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	// Extract Members
	ms := []plugs.Member{}
	for _, kv := range res.Kvs {
		m := plugs.Member{}
		m.Name = p.scanKey(kv.Key)
		m.Host, m.Port = p.scanValue(kv.Value)
		ms = append(ms, m)
	}

	return ms, nil
}

// Watch implements plugs.Plug.
func (p *etcdPlug) Watch(ctx context.Context, cname string, size int) <-chan plugs.Change {
	p.wmx.Lock()
	defer p.wmx.Unlock()

	if _, ok := p.w[cname]; ok {
		// There's already a watch channel
		return p.w[cname]
	}

	watch := p.client.Watch(ctx, p.keyspace(cname), clientv3.WithPrefix())
	stopw := make(chan bool)
	w := make(chan plugs.Change, size)

	p.w[cname] = w
	p.stopw[cname] = stopw

	go func() {
		for {
			select {
			case res, ok := <- watch:
				if !ok {
					// Something went wrong.
					go p.Stop(cname)
					return
				}
				for _, event := range res.Events {
					switch event.Type {
					case clientv3.EventTypePut:
						m := plugs.Member{}
						m.Name = p.scanKey(event.Kv.Key)
						m.Host, m.Port = p.scanValue(event.Kv.Value)
						
						p.memx.RLock()
						if p.me != nil && p.me.Name == m.Name {
							p.memx.RUnlock()
							continue
						}
						p.memx.RUnlock()

						w <- plugs.Change{
							Method: plugs.MethodPut,
							Data: m,
						}
					case clientv3.EventTypeDelete:
						m := plugs.Member{
							Name: p.scanKey(event.Kv.Key),
						}

						p.memx.RLock()
						if p.me != nil && p.me.Name == m.Name {
							p.memx.RUnlock()
							continue
						}
						p.memx.RUnlock()

						w <- plugs.Change{
							Method: plugs.MethodDel,
							Data: m,
						}
					}
				}
			case <- stopw:
				// Stop watching changes
				return
			}
		}
	} ()

	return w
}

// Stop implements plugs.Plug.
func (p *etcdPlug) Stop(cname string) {
	p.wmx.Lock()
	defer p.wmx.Unlock()
	
	if _, ok := p.w[cname]; ok {
		p.stopw[cname] <- true
		delete(p.w, cname)
		delete(p.stopw, cname)
	}
}

// Destroy implements plugs.Plug.
func (p *etcdPlug) Destroy() {
	p.wmx.Lock()

	for _, v := range p.stopw {
		v <- true
	}

	p.client.Close()
}

func (p *etcdPlug) keyspace(cname string) string {
	return fmt.Sprintf(fKeyspace, p.prefix, cname)
}

func (p *etcdPlug) key(cname, name string) string {
	return fmt.Sprintf(fKey, p.prefix, cname, name)
}

func (p *etcdPlug) value(host, port string) string {
	return fmt.Sprintf(fValue, host, port)
}

func (p *etcdPlug) scanKey(key []byte) string {
	for i := len(key) - 1; i >= 0; i-- {
		if key[i] == ':' {
			return string(key[i+1:])
		}
	}
	return ""
}

func (p *etcdPlug) scanValue(value []byte) (string, string) {
	for i := 0; i < len(value); i++ {
		if value[i] == ':' {
			return string(value[:i]), string(value[i+1:])
		}
	}
	return "", ""
}