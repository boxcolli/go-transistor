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
	Bus(w io.StreamWriter) (io.Bus, bool)	// 
	Work(w io.StreamWriter) error
	Stop(w io.StreamWriter)
	StopAll()
}
