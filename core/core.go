package core

import (
	"context"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

// A core is a builder and also a CLI engine
type Core interface {
	Collect(r io.StreamReader) error			// Block function
	Emit(w io.StreamWriter) error				// Block function
	
	Apply(w io.StreamWriter, cg *types.Change)
	Delete(w io.StreamWriter)

	Command(ctx context.Context, args []string) (<-chan string, error)
}
