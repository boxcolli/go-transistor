package emitters

import (
	"github.com/boxcolli/go-transistor/types"
)

// Emitter receives new messages and emits through outlet.
type Emitter interface {
	Start() <-chan error
	Stop()
	Emit(m *types.Message)	// goroutine safe; use channel to implement queue
}
