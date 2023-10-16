package redisplug

import (
	"context"
	"strings"
	"sync"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
	"github.com/redis/go-redis/v9"
)

type redisPlug struct {
	client *redis.Client
	f      plugs.Formatter

	// To prevent emitting myself in watch channel
	me   *types.Member
	memx sync.RWMutex

	ch   map[string]chan *plugs.Event // Singleton watch channels
	stop map[string]chan bool         // Channels connected with watch goroutines
	chmx sync.Mutex
}

/*
	redis follows the format on keyspace notification:
		__keyspace@<db>__:<key>
	Use NewBasicRedisFormatter() or customized formatter for redis.
*/
func NewRedisPlug(client *redis.Client, f plugs.Formatter) plugs.Plug {
	return &redisPlug{
		client: client,
		f: f,
		me: nil,
		memx: sync.RWMutex{},
		ch: make(map[string]chan *plugs.Event),
		stop: make(map[string]chan bool),
		chmx: sync.Mutex{},
	}
}

const (
	eventSet = "set"
	eventDel = "del"
	eventExp = "expire"
)

// Me implements plugs.Plug.
func (p *redisPlug) Me(ctx context.Context, op types.Operation, me *types.Member) error {
	p.memx.Lock()
	defer p.memx.Unlock()

	if op == types.OperationAdd {
		// Cache myself
		p.me = me

		// Advertise myself
		err := p.client.Set(ctx, p.f.PrintKey(me), p.f.PrintValue(me), 0).Err()
		if err != nil {
			return err
		}

	} else if op == types.OperationDel {
		err := p.client.Del(ctx, p.f.PrintKey(p.me)).Err()
		if err != nil {
			return err
		}

		p.me = nil
	}

	return nil
}

// Watch implements plugs.Plug.
func (p *redisPlug) Watch(ctx context.Context, cname string, size int) (<-chan *plugs.Event, error) {
	p.chmx.Lock()
	defer p.chmx.Unlock()

	if ch, ok := p.ch[cname]; ok {
		// There is already a watch channel
		return ch, nil
	}

	var watch <-chan *redis.Message
	{
		pubsub := p.client.PSubscribe(context.Background(), p.f.PrintKeyspace(cname))
		watch = pubsub.Channel()
	}
	ch := make(chan *plugs.Event, size)
	stop := make(chan bool)

	// Before issuing new events, push already existing KVs.
	{
		res, err := p.client.Keys(ctx, p.f.PrintKeyspace(cname)).Result()
		if err != nil {
			return nil, err
		}

		for _, key := range res {
			value, err := p.client.Get(ctx, key).Result()
			if err != nil {
				return nil, err
			}
			e := new(plugs.Event)
			e.Op = types.OperationAdd
			p.f.ScanKey(key, &e.Data)
			p.f.ScanValue(value, &e.Data)
			ch <- e
		}
	}

	go func(watch <-chan *redis.Message, ch chan<- *plugs.Event, stop <-chan bool) {
		for {
			select {
				// Stop watching changes
			case <- stop:
				return

			case msg, ok := <- watch:
				if !ok {
					// Something went wrong.
					defer p.Stop(cname)
					return
				}

				key := strings.SplitN(msg.Channel, ":", 2)[1]

				// Fetch events
				switch msg.Payload {
				case eventSet:
					value, err := p.client.Get(ctx, key).Result()
					if err != nil {
						// ignore error
						continue
					}

					e := new(plugs.Event)
					e.Op = types.OperationAdd
					p.f.ScanKey(key, &e.Data)
					p.f.ScanValue(value, &e.Data)

					// Discard if it's me
					p.memx.RLock()
					if p.me != nil &&
						p.me.Cname == e.Data.Cname &&
						p.me.Name == e.Data.Name {
						p.memx.RUnlock()
						continue
					}
					p.memx.RUnlock()

					ch <- e

				case eventDel:
					e := new(plugs.Event)
					e.Op = types.OperationDel
					p.f.ScanKey(key, &e.Data)

					// Discard if it's me
					p.memx.RLock()
					if p.me != nil &&
						p.me.Cname == e.Data.Cname &&
						p.me.Name == e.Data.Name {
						p.memx.RUnlock()
						continue
					}
					p.memx.RUnlock()

					ch <- e
				
				case eventExp:	// pretend it's delete
					e := new(plugs.Event)
					e.Op = types.OperationDel
					p.f.ScanKey(key, &e.Data)

					// Discard if it's me
					p.memx.RLock()
					if p.me != nil &&
						p.me.Cname == e.Data.Cname &&
						p.me.Name == e.Data.Name {
						p.memx.RUnlock()
						continue
					}
					p.memx.RUnlock()

					ch <- e
				}
			}
		}
	} (watch, ch, stop)

	p.ch[cname] = ch
	p.stop[cname] = stop

	return ch, nil
}

// Stop implements plugs.Plug.
func (p *redisPlug) Stop(cname string) {
	p.chmx.Lock()
	defer p.chmx.Unlock()
	
	if _, ok := p.ch[cname]; ok {
		p.stop[cname] <- true
		delete(p.ch, cname)
		delete(p.stop, cname)
	}
}

// Close implements plugs.Plug.
func (p *redisPlug) Close() {
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
