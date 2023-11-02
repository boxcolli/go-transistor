package emitter

import (
	"errors"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrClosed = errors.New("streamwrite is closed")
)

// Emitter receives new messages and emits through outlet.
type Emitter interface {
	Work(w io.StreamWriter) error
	Stop()
	Emit(m *types.Message)	// should be goroutine safe
}
