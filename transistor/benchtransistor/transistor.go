package benchtransistor

import (
	"context"

	"github.com/boxcolli/go-transistor/transistor"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/middleware"
	"github.com/boxcolli/go-transistor/types"
)

type Option struct {
}

type benchTransistor struct {
	com transistor.Component
	opt Option
	mid middleware.BaseMiddleware
}

func NewBenchTransistor(com transistor.Component, opt Option) transistor.Transistor {
	return &benchTransistor{
		com: com,
		opt: opt,
		mid: middleware.NewBaseMiddleware(com.Base),
	}
}

func (c *benchTransistor) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.mid, r)
}
func (c *benchTransistor) Emit(w io.StreamWriter) error {
	return c.com.Emitter.Work(w)
}
func (c *benchTransistor) Apply(w io.StreamWriter, cg *types.Change) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		c.mid.Apply(bus, cg)
		return nil
	} else {
		return transistor.ErrNotFound
	}
}
func (c *benchTransistor) Stop(w io.StreamWriter) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		bus.Lock()
		c.mid.Delete(bus)
		c.com.Emitter.Stop(w)
		return nil
	} else {
		return transistor.ErrNotFound
	}
}
func (c *benchTransistor) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
