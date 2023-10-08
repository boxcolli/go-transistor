package sinks

import (
	"github.com/boxcolli/go-transistor/types"
)

//
type Sink interface {
	Create(types.Message) error
	Update(types.Message) error
	Delete(types.Message) error
}

type SinkOption struct {
	// validateTopicTable	bool
}
