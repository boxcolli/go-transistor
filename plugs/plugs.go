package plugs

import "context"

const (
	MethodPut = byte(0)
	MethodDel = byte(1)
)

/*
	Plug is a wrapper of a database client which is connected with Cluster member store.
	Plug is a singleton and handles all the clusters related to the node.
	Plug doesn't handle connection between nodes.
*/
type Plug interface {
	// Advertise myself as a publisher
	Advertise(ctx context.Context, cname string, me Member) error

	// Read members of the cluster
	Read(ctx context.Context, cname string) ([]Member, error)
	
	// Returns a channel that emits changes on cluster publisher.
	// Each channel must be a singleton.
	Watch(ctx context.Context, cname string, size int) <-chan Change

	// Stop watching changes on a Cluster
	Stop(cname string)

	// Stop all watching changes and close client 
	Destroy()
}

type Change struct {
	Method	byte
	Data	Member
}

type Member struct {
	Cname string
	Name string
	Host string
	Port string
}
