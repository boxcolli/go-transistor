package emitters

import (
	"github.com/boxcolli/go-transistor/types"
)

type Emitter interface {
	Start() <-chan error
	Stop()
	Emit(m *types.Message)
}
