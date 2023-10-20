package dialers

import (
	"context"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

// Dialer takes care for grpc clients and their connections.
type Dialer interface {
	Dial(ctx context.Context, m *types.Member) (io.StreamReader, error)
	Close(m *types.Member) error
	CloseAll()
}
