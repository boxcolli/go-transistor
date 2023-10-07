package streamemitter

import (
	"github.com/boxcolli/go-transistor/emitters"
	"github.com/boxcolli/go-transistor/types"
	"google.golang.org/grpc"
)

type streamEmitter struct {
}

// Emit implements emitters.Emitter.
func (*streamEmitter) Emit(m *types.Message) {
	panic("unimplemented")
}

// Start implements emitters.Emitter.
func (*streamEmitter) Start() <-chan error {
	panic("unimplemented")
}

// Stop implements emitters.Emitter.
func (*streamEmitter) Stop() {
	panic("unimplemented")
}

func NewStreamEmitter(s grpc.ServerStream) emitters.Emitter {
	return &streamEmitter{}
}
