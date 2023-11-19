package simplecore

import (
	"context"

	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type Option struct {
}

type simpleCore struct {
	com core.Component
	opt Option
}

func NewSimpleCore(com core.Component, opt Option) core.Core {
	return &simpleCore{
		com: com,
		opt: opt,
	}
}

func (c *simpleCore) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.com.Base, r)
}
func (c *simpleCore) Emit(w io.StreamWriter) error {
	return c.com.Emitter.Work(w)
}
func (c *simpleCore) Apply(w io.StreamWriter, cg *types.Change) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		c.com.Base.Apply(bus, cg)
		return nil
	} else {
		return core.ErrNotFound
	}
}
func (c *simpleCore) Stop(w io.StreamWriter) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		bus.Lock()
		c.com.Base.Delete(bus)
		c.com.Emitter.Stop(w)
		return nil
	} else {
		return core.ErrNotFound
	}
}
func (c *simpleCore) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
