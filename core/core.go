package core

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

// A core is a builder and also a CLI engine
type Core interface {
	Start()

	// Block functions
	Collect(io.StreamReader) error
	// Emit(emitter.Emitter, io.StreamWriter) error // further implementation
	Apply(emitter.Emitter, *types.Change)
	Delete(emitter.Emitter)
	Command(args []string) <-chan string
}
