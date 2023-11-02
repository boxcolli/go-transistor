package base

import (
	"errors"

	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrNoTopic = errors.New("topic is empty")
)

type Base interface {
	Start()
	Stop()
	Flow(m *types.Message)
	Apply(e emitter.Emitter, cg *types.Change)
	Delete(e emitter.Emitter)
}
