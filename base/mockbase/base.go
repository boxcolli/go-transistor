package mockbase

import (
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type mockBase struct {
}

// Apply implements base.Base.
func (*mockBase) Apply(e *emitter.Emitter, op types.Operation, topics []types.Topic) {
	panic("unimplemented")
}

// Flow implements base.Base.
func (*mockBase) Flow(m *types.Message) error {
	panic("unimplemented")
}

func NewMockBase() base.Base {
	return &mockBase{}
}
