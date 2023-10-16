package etcdplug

import (
	"context"
	"sync"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
	"go.etcd.io/etcd/client/v3"
)

type etcdPlug struct {
	client *clientv3.Client
	f plugs.Formatter

	// To prevent emitting myself in watch channel
	me   *types.Member
	memx sync.RWMutex

	ch   map[string]chan *plugs.Event	// Singleton watch channels
	stop map[string]chan bool         	// Channels connected with watch goroutines
	chmx sync.Mutex
}

func NewEtcdPlug(cl *clientv3.Client, f plugs.Formatter) plugs.Plug {
	return &etcdPlug{
		client: cl,
		f: f,
		me: nil,
		memx: sync.RWMutex{},
		ch: map[string]chan *plugs.Event{},
		stop: map[string]chan bool{},
		chmx: sync.Mutex{},
	}
}

// Advertise implements plugs.Plug.
func (p *etcdPlug) Me(ctx context.Context, op types.Operation, me *types.Member) error {
	p.memx.Lock()
	defer p.memx.Unlock()
	
	if op == types.OperationAdd {
		// Write myself on local cache
		p.me = me
		
		// Advertise myself on the KV store
		_, err := p.client.Put(ctx, p.f.PrintKey(me), p.f.PrintValue(me))
		if err != nil {
			return err
		}

	} else if op == types.OperationDel {
		_, err := p.client.Delete(ctx, p.f.PrintKey(p.me))
		if err != nil {
			return err
		}

		p.me = nil
	}

	return nil
}

// Watch implements plugs.Plug.
func (p *etcdPlug) Watch(ctx context.Context, cname string, size int) (<-chan *plugs.Event, error) {
	p.chmx.Lock()
	defer p.chmx.Unlock()

	if ch, ok := p.ch[cname]; ok {
		// There's already a watch channel
		return ch, nil
	}

	watch := p.client.Watch(ctx, p.f.PrintKeyspace(cname), clientv3.WithPrefix())
	ch := make(chan *plugs.Event, size)
	stop := make(chan bool)

	// Before issuing new events from watch,
	// the existing key-value pairs should be sent first.
	{
		res, err := p.client.Get(ctx, p.f.PrintKeyspace(cname), clientv3.WithPrefix())
		if err != nil {
			return nil, err
		}

		for _, kv := range res.Kvs {
			e := new(plugs.Event)
			e.Op = types.OperationAdd
			p.f.ScanKey(string(kv.Key), &e.Data)
			p.f.ScanValue(string(kv.Value), &e.Data)
			ch <- e
		}
	}

	go func(watch clientv3.WatchChan, ch chan<- *plugs.Event, stop <-chan bool) {
		for {
			select {
			case <- stop:
				// Stop watching changes
				return

			case res, ok := <- watch:
				if !ok {
					// Something went wrong.
					defer p.Stop(cname)
					return
				}

				// Fetch events
				for _, event := range res.Events {
					switch event.Type {
					case clientv3.EventTypePut:
						// Scan
						e := new(plugs.Event)
						e.Op = types.OperationAdd
						p.f.ScanKey(string(event.Kv.Key), &e.Data)
						p.f.ScanValue(string(event.Kv.Value), &e.Data)
						
						// Discard if it's me
						p.memx.RLock()
						if p.me != nil &&
							p.me.Cname == e.Data.Cname &&
							p.me.Name == e.Data.Name {
							p.memx.RUnlock()
							continue
						}
						p.memx.RUnlock()

						// Send event
						ch <- e

					case clientv3.EventTypeDelete:
						// Scan
						e := new(plugs.Event)
						e.Op = types.OperationDel
						p.f.ScanKey(string(event.Kv.Key), &e.Data)

						// Discard if its me
						p.memx.RLock()
						if p.me != nil &&
							p.me.Cname == e.Data.Cname &&
							p.me.Name == e.Data.Name {
							p.memx.RUnlock()
							continue
						}
						p.memx.RUnlock()

						// Send event
						ch <- e
					}
				}
			}
		}
	} (watch, ch, stop)

	p.ch[cname] = ch
	p.stop[cname] = stop

	return ch, nil
}

// Stop implements plugs.Plug.
func (p *etcdPlug) Stop(cname string) {
	p.chmx.Lock()
	defer p.chmx.Unlock()
	
	if _, ok := p.ch[cname]; ok {
		p.stop[cname] <- true
		delete(p.ch, cname)
		delete(p.stop, cname)
	}
}

// Destroy implements plugs.Plug.
func (p *etcdPlug) Close() {
	p.chmx.Lock()
	defer p.chmx.Unlock()

	for _, v := range p.stop {
		v <- true
	}
	for _, v := range p.ch {
		close(v)
	}
	p.client.Close()
	p.ch = make(map[string]chan *plugs.Event)
	p.stop = make(map[string]chan bool)
}
