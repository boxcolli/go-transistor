package benchcore

import (
	"context"

	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/middleware"
	"github.com/boxcolli/go-transistor/types"
)

type Option struct {
}

type benchCore struct {
	com core.Component
	opt Option
	mid middleware.BaseMiddleware
}

func NewBenchCore(com core.Component, opt Option) core.Core {
	return &benchCore{
		com: com,
		opt: opt,
		mid: middleware.NewBaseMiddleware(com.Base),
	}
}

func (c *benchCore) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.mid, r)
}
func (c *benchCore) Emit(w io.StreamWriter) error {
	return c.com.Emitter.Work(w)
}
func (c *benchCore) Apply(w io.StreamWriter, cg *types.Change) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		c.mid.Apply(bus, cg)
		return nil
	} else {
		return core.ErrNotFound
	}
}
func (c *benchCore) Stop(w io.StreamWriter) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		bus.Lock()
		c.mid.Delete(bus)
		c.com.Emitter.Stop(w)
		return nil
	} else {
		return core.ErrNotFound
	}
}
func (c *benchCore) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
