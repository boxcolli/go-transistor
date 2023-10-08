package base

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type Base interface {
	Flow(m *types.Message) error
	Apply(e *emitter.Emitter, f *types.Filter)
	Delete(e *emitter.Emitter)
}
