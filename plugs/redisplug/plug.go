package redisplug

import (
	"context"
	"sync"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
	"github.com/redis/go-redis/v9"
)

type redisPlug struct {
	client *redis.Client

	// Special key prefix for key name
	prefix string

	// To prevent emitting myself in watch channel
	me   *types.Member
	memx sync.RWMutex

	w     map[string]chan plugs.Change // Singleton watch channels
	stopw map[string]chan bool         // Channels connected with watch goroutines
	wmx   sync.Mutex
}

// Advertise implements plugs.Plug.
func (*redisPlug) Advertise(ctx context.Context, cname string, me plugs.Member) error {
	panic("unimplemented")
}

// Destroy implements plugs.Plug.
func (*redisPlug) Destroy() {
	panic("unimplemented")
}

// Read implements plugs.Plug.
func (*redisPlug) Read(ctx context.Context, cname string) ([]plugs.Member, error) {
	panic("unimplemented")
}

// Stop implements plugs.Plug.
func (*redisPlug) Stop(cname string) {
	panic("unimplemented")
}

// Watch implements plugs.Plug.
func (*redisPlug) Watch(ctx context.Context, cname string, size int) <-chan plugs.Change {
	panic("unimplemented")
}

func NewRedisPlug(client *redis.Client, prefix string) plugs.Plug {
	return &redisPlug{
		client: client,
	}
}
