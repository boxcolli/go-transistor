package sinkemitter

import (
	"github.com/boxcolli/go-transistor/emitters"
	"github.com/boxcolli/go-transistor/sinks"
	"github.com/boxcolli/go-transistor/types"
)

type sinkEmitter struct {
}

// Emit implements emitters.Emitter.
func (*sinkEmitter) Emit(m *types.Message) {
	panic("unimplemented")
}

// Start implements emitters.Emitter.
func (*sinkEmitter) Start() <-chan error {
	panic("unimplemented")
}

// Stop implements emitters.Emitter.
func (*sinkEmitter) Stop() {
	panic("unimplemented")
}

func NewSinkEmitter(s *sinks.Sink) emitters.Emitter {
	return &sinkEmitter{}
}
