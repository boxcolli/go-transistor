package simpletransistor

import (
	"context"

	"github.com/boxcolli/go-transistor/transistor"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type Option struct {
}

type simpleTransistor struct {
	com transistor.Component
	opt Option
}

func NewSimpleCore(com transistor.Component, opt Option) transistor.Transistor {
	return &simpleTransistor{
		com: com,
		opt: opt,
	}
}

func (c *simpleTransistor) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.com.Base, r)
}
func (c *simpleTransistor) Emit(w io.StreamWriter) error {
	return c.com.Emitter.Work(w)
}
func (c *simpleTransistor) Apply(w io.StreamWriter, cg *types.Change) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		c.com.Base.Apply(bus, cg)
		return nil
	} else {
		return transistor.ErrNotFound
	}
}
func (c *simpleTransistor) Stop(w io.StreamWriter) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		bus.Lock()
		c.com.Base.Delete(bus)
		c.com.Emitter.Stop(w)
		return nil
	} else {
		return transistor.ErrNotFound
	}
}
func (c *simpleTransistor) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
