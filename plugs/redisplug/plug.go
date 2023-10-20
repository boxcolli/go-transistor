package redisplug

import (
	"context"
	"strings"
	"sync"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
	"github.com/redis/go-redis/v9"
)

type entry struct {
	ch		chan *plugs.Event	// Singleton watch channels
	stop	chan bool			// Channels connected with watch goroutines
}

type redisPlug struct {
	client *redis.Client
	f      RedisFormatter

	// To prevent emitting myself in watch channel
	me   *types.Member
	memx sync.RWMutex

	// Per cluster entries
	e   map[string]entry
	emx sync.Mutex
}

/*
	redis follows the format on keyspace notification:
		__keyspace@<db>__:<key>
	Use NewBasicRedisFormatter() or customized formatter for redis.
*/
func NewRedisPlug(client *redis.Client, f RedisFormatter) plugs.Plug {
	return &redisPlug{
		client: client,
		f: f,
		me: nil,
		memx: sync.RWMutex{},
		e: make(map[string]entry),
		emx: sync.Mutex{},
	}
}

const (
	eventSet = "set"
	eventDel = "del"
	eventExp = "expired"
)

// Me implements plug.Plug.
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

// Watch implements plug.Plug.
func (p *redisPlug) Watch(ctx context.Context, cname string, size int) (<-chan *plugs.Event, error) {
	p.emx.Lock()
	defer p.emx.Unlock()

	if ent, ok := p.e[cname]; ok {
		// There is already a watch channel
		return ent.ch, nil
	}

	var watch	<-chan *redis.Message
	var ch		chan *plugs.Event
	var stop 	chan bool
	{
		pubsub := p.client.PSubscribe(context.Background(), p.f.PrintPSubscribeKeyspace(cname))
		watch = pubsub.Channel()

		ch = make(chan *plugs.Event, size)
		stop = make(chan bool)

		p.e[cname] = entry{ ch: ch, stop: stop }
	}

	// Before issuing new events, push already existing KVs.
	{
		res, err := p.client.Keys(ctx, p.f.PrintKeyspace(cname)).Result()
		if err != nil {
			return nil, err
		}

		p.memx.RLock()
		me := p.me
		p.memx.RUnlock()

		for _, key := range res {
			value, err := p.client.Get(ctx, key).Result()
			if err != nil {
				return nil, err
			}
			e := new(plugs.Event)
			e.Op = types.OperationAdd
			p.f.ScanKey(key, &e.Data)
			p.f.ScanValue(value, &e.Data)

			// Skip if it's me
			if me != nil && me.EqualsId(e.Data) {
				continue
			}
			
			ch <- e
		}
	}

	go p.watch(ctx, cname, watch, ch, stop)

	return ch, nil
}

func (p *redisPlug) watch(ctx context.Context, cname string, watch <-chan *redis.Message, ch chan<- *plugs.Event, stop <-chan bool) {
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
			e := new(plugs.Event)
			{
				p.f.ScanKey(key, &e.Data)

				p.memx.RLock()
				me := p.me
				p.memx.RUnlock()

				// Discard if it's me
				if me != nil && me.EqualsId(e.Data) {
					continue
				}
			}

			// Fetch events
			switch msg.Payload {
			case eventSet:
				value, err := p.client.Get(ctx, key).Result()
				if err != nil {
					// ignore error
					continue
				}

				e.Op = types.OperationAdd
				p.f.ScanValue(value, &e.Data)

				ch <- e

			case eventDel:
				e.Op = types.OperationDel
				p.f.ScanKey(key, &e.Data)

				ch <- e
			
			case eventExp:	// pretend it's delete
				e := new(plugs.Event)
				e.Op = types.OperationDel
				p.f.ScanKey(key, &e.Data)

				ch <- e
			}
		}
	}
}

// Stop implements plugs.Plug.
func (p *redisPlug) Stop(cname string) {
	p.emx.Lock()
	defer p.emx.Unlock()
	
	if ent, ok := p.e[cname]; ok {
		close(ent.ch)
		close(ent.stop)
		delete(p.e, cname)
	}
}

// Close implements plugs.Plug.
func (p *redisPlug) Close() {
	p.emx.Lock()
	defer p.emx.Unlock()

	for _, ent := range p.e {
		close(ent.stop)
		close(ent.ch)
	}

	p.client.Close()
	p.e = make(map[string]entry)
}
