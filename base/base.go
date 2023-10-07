package base

import (
	"github.com/boxcolli/go-transistor/emitters"
	"github.com/boxcolli/go-transistor/types"
)

type Base interface {
	Load(f *types.Filter, e *emitters.Emitter)
	Unload(e *emitters.Emitter)
	Flow(m *types.Message)
	Apply(newF *types.Filter, e *emitters.Emitter)
}