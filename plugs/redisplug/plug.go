package redisplug

import (
	"sync"

	"github.com/boxcolli/pepperlink/plugs"
	"github.com/redis/go-redis/v9"
)

type redisPlug struct {
	client *redis.Client

	gk func(string, string) string
	gv func(string, string) string

	// To prevent emitting myself in watch channel
	me   *plugs.Member
	memx sync.RWMutex

	w   map[string]chan plugs.Change // Singleton watch channels
	wch map[string]chan bool         // Channels connected with watch goroutines
	wmx sync.Mutex
}

// generateKey implements plugs.Plug.
func (p *redisPlug) generateKey(cname string, name string) string {
	return p.gk(cname, name)
}

// generateValue implements plugs.Plug.
func (p *redisPlug) generateValue(host string, port string) string {
	return p.gv(host, port)
}

// Advertise implements plugs.Plug.
func (p *redisPlug) Advertise(cname string, me plugs.Member) {

}

// Destroy implements plugs.Plug.
func (p *redisPlug) Destroy() {
	panic("unimplemented")
}

// Read implements plugs.Plug.
func (p *redisPlug) Read(cname string) []plugs.Member {
	panic("unimplemented")
}

// Stop implements plugs.Plug.
func (p *redisPlug) Stop(cname string) {
	panic("unimplemented")
}

// Watch implements plugs.Plug.
func (p *redisPlug) Watch(cname string) <-chan plugs.Change {
	panic("unimplemented")
}

func NewRedisPlug(client *redis.Client, gk func(string, string) string, gv func(string, string) string) plugs.Plug {
	return &redisPlug{
		client: client,
	}
}
