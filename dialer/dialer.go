package dialer

import (
	"github.com/boxcolli/go-transistor/types"
	"github.com/boxcolli/go-transistor/io"
)

// Dialer takes care for grpc clients and their connections.
type Dialer interface {
	Dial(m *types.Member) (io.StreamReader, error)
	Close(m *types.Member) error
	CloseAll()
}
