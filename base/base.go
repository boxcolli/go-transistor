package base

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type Base interface {
	Start()
	Stop()
	Flow(m *types.Message) error
	Apply(e *emitter.Emitter, op types.Operation, topics []types.Topic)
	Delete(e *emitter.Emitter)
}
 