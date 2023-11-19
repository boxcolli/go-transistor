package core

import (
	"context"
	"errors"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

var (
	ErrNotFound = errors.New("entry not found")
)

// A core is a builder and also a CLI engine
type Core interface {
	// Block functions
	Collect(r io.StreamReader) error
	Emit(w io.StreamWriter) error
	Apply(w io.StreamWriter, cg *types.Change) error
	Stop(w io.StreamWriter) error

	Command(ctx context.Context, args []string) (<-chan string, error)
}

type Component struct {
	Collector 	collector.Collector
	Base      	base.Base
	Emitter		emitter.Emitter
}
