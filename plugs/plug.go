package plugs

import (
	"context"

	"github.com/boxcolli/go-transistor/types"
)

/*
	Plug is a wrapper of a database client which is connected with Cluster member store.
	Plug is a singleton and handles all the clusters and nodes related to the node.
	Plug doesn't handle connection between nodes.
*/
type Plug interface {
	// Advertise myself as a publisher
	Me(ctx context.Context, op types.Operation, me *types.Member) error

	// Returns a channel that emits changes on cluster publisher.
	// Each channel must be a singleton.
	Watch(ctx context.Context, cname string, size int) (<-chan *Event, error)

	// Stop watching changes on a Cluster
	Stop(cname string)

	// Stop all watching changes and close client 
	Close()
}
