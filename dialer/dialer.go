package dialer

import (
	"context"
	"errors"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrMemberNotFound = errors.New("the member is not found")
)

// Dialer takes care for grpc clients and their connections.
type Dialer interface {
	Dial(ctx context.Context, m *types.Member, c *types.Change) (io.StreamReader, error)
	Apply(m *types.Member, c *types.Change) error
	Close(m *types.Member) error
	CloseAll()
}
