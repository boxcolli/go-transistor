package dialer

import "github.com/boxcolli/go-transistor/types"

type Connecter interface {
	Connect(Address) StreamReader
	Handle(e chan error)
}

// Dialer takes care for grpc clients and their connections.
type  interface {

	// 1. 연결을 만든다
	// 2. go Collector.Work()
	Dial(m types.Member)
	Add(m types.Member)
	Delete(m types.Member)
}
