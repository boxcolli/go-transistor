package etcdplug

import (
	"context"
	"sync"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
	"go.etcd.io/etcd/client/v3"
)

type entry struct {
	ch		chan *plugs.Event	// Singleton watch channels
	stop	chan bool			// Channels connected with watch goroutines
}

type etcdPlug struct {
	client *clientv3.Client
	f plugs.Formatter

	// To prevent emitting myself in watch channel
	me   *types.Member
	memx sync.RWMutex

	// Per cluster entries
	e   map[string]entry
	emx sync.Mutex
}

func NewEtcdPlug(cl *clientv3.Client, f plugs.Formatter) plugs.Plug {
	return &etcdPlug{
		client: cl,
		f: f,
		me: nil,
		memx: sync.RWMutex{},
		e: make(map[string]entry),
		emx: sync.Mutex{},
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
	p.emx.Lock()
	defer p.emx.Unlock()

	if ent, ok := p.e[cname]; ok {
		// There is already a watch channel
		return ent.ch, nil
	}

	var watch	clientv3.WatchChan
	var ch		chan *plugs.Event
	var stop 	chan bool
	{
		watch = p.client.Watch(ctx, p.f.PrintKeyspace(cname), clientv3.WithPrefix())

		ch = make(chan *plugs.Event, size)
		stop = make(chan bool)

		p.e[cname] = entry{ ch: ch, stop: stop }
	}

	// Before issuing new events, push already existing KVs.
	{
		res, err := p.client.Get(ctx, p.f.PrintKeyspace(cname), clientv3.WithPrefix())
		if err != nil {
			return nil, err
		}

		p.memx.RLock()
		me := p.me
		p.memx.RUnlock()

		for _, kv := range res.Kvs {
			e := new(plugs.Event)
			e.Op = types.OperationAdd
			p.f.ScanKey(string(kv.Key), &e.Data)
			p.f.ScanValue(string(kv.Value), &e.Data)

			// Skip if it's me
			if me != nil && me.EqualsId(e.Data) {
				continue
			}

			ch <- e
		}
	}

	go p.watch(cname, watch, ch, stop)

	return ch, nil
}

func (p *etcdPlug) watch(cname string, watch clientv3.WatchChan, ch chan<- *plugs.Event, stop <-chan bool) {
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
				e := new(plugs.Event)
				{
					p.f.ScanKey(string(event.Kv.Key), &e.Data)
					// Discard if it's me
					p.memx.RLock()
					me := p.me
					p.memx.RUnlock()

					if me != nil && me.EqualsId(e.Data) {
						continue
					}
				}

				switch event.Type {
				case clientv3.EventTypePut:
					e.Op = types.OperationAdd
					p.f.ScanValue(string(event.Kv.Value), &e.Data)

					// Send event
					ch <- e

				case clientv3.EventTypeDelete:
					e.Op = types.OperationDel

					// Send event
					ch <- e
				}
			}
		}
	}
}

// Stop implements plugs.Plug.
func (p *etcdPlug) Stop(cname string) {
	p.emx.Lock()
	defer p.emx.Unlock()
	
	if ent, ok := p.e[cname]; ok {
		close(ent.ch)
		close(ent.stop)
		delete(p.e, cname)
	}
}

// Close implements plugs.Plug.
func (p *etcdPlug) Close() {
	p.emx.Lock()
	defer p.emx.Unlock()

	for _, ent := range p.e {
		close(ent.stop)
		close(ent.ch)
	}

	p.client.Close()
	p.e = make(map[string]entry)
}
