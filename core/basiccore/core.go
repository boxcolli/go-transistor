package basiccore

import (
	"context"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/core"
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

type basicCore struct {
	com Component
	opt Option
	mid middleware.BaseMiddleware
}

func NewBasicCore(com Component, opt Option) core.Core {
	return &basicCore{
		com: com,
		opt: opt,
		mid: middleware.NewBaseMiddleware(com.Base),
	}
}

func (c *basicCore) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.mid, r)
}
func (c *basicCore) Emit(w io.StreamWriter) error {
	c.com.Emitter.Bus(w)
	return c.com.Emitter.Work(w)
}

func (c *basicCore) Apply(w io.StreamWriter, cg *types.Change) {
	bus, _ := c.com.Emitter.Bus(w)
	c.mid.Apply(bus, cg)
}
func (c *basicCore) Delete(w io.StreamWriter) {
	bus, _ := c.com.Emitter.Bus(w)
	c.mid.Delete(bus)
	c.com.Emitter.Stop(w)
}

func (c *basicCore) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
