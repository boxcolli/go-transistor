package basictransistor

import (
	"context"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/transistor"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/middleware"
	"github.com/boxcolli/go-transistor/types"
)

type Component struct {
	Collector 	collector.Collector
	Base      	base.Base
	Emitter		emitter.Emitter
}

type Option struct {
}

type basicTransistor struct {
	com Component
	opt Option
	mid middleware.BaseMiddleware
}

func NewBasicCore(com Component, opt Option) transistor.Transistor {
	return &basicTransistor{
		com: com,
		opt: opt,
		mid: middleware.NewBaseMiddleware(com.Base),
	}
}

func (c *basicTransistor) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.mid, r)
}
func (c *basicTransistor) Emit(w io.StreamWriter) error {
	return c.com.Emitter.Work(w)
}
func (c *basicTransistor) Apply(w io.StreamWriter, cg *types.Change) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		c.mid.Apply(bus, cg)
		return nil
	} else {
		return transistor.ErrNotFound
	}
}
func (c *basicTransistor) Stop(w io.StreamWriter) error {
	if bus, ok := c.com.Emitter.Bus(w); ok {
		bus.Lock()
		c.mid.Delete(bus)
		c.com.Emitter.Stop(w)
		return nil
	} else {
		return transistor.ErrNotFound
	}
}
func (c *basicTransistor) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
