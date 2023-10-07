package plugs

import (
	"context"

	"github.com/boxcolli/go-transistor/types"
)

type Change struct {
	Method	types.Method
	Data	types.Member
}

type Client interface {
	Watch(ctx context.Context, cname string, size int) <-chan Change
	Close()
}

