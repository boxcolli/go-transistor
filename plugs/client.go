package plugs

import (
	"context"

	"github.com/boxcolli/go-transistor/types"
)

type Method int

const (
	EventPut Method = iota
	EventDel
)

type Event struct {
	Method	Method
	Data	types.Member
}

type Client interface {
	Watch(ctx context.Context, cname string, size int) <-chan *Event
	Close()
}
