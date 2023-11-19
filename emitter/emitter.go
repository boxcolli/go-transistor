package emitter

import (
	"errors"

	"github.com/boxcolli/go-transistor/io"
)

var (
	ErrClosed = errors.New("streamwrite is closed")
)

// Emitter receives new messages and emits through outlet.
type Emitter interface {
	Work(w io.StreamWriter) error
	Bus(w io.StreamWriter) (io.Bus, bool)
	Stop(w io.StreamWriter)
	StopAll()
}
