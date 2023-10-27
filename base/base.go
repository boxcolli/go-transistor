package base

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type Base interface {
	Start()
	Stop()
	Flow(m *types.Message) error
	Apply(e *emitter.Emitter, cg *types.Change)
	Delete(e *emitter.Emitter)
}
