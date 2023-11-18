package simplecore

import (
	"context"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type Component struct {
	Collector 	collector.Collector
	Base      	base.Base
	Emitter		emitter.Emitter
}

type Option struct {
}

type simpleCore struct {
	com Component
	opt Option
}

func NewsimpleCore(com Component, opt Option) core.Core {
	return &simpleCore{
		com: com,
		opt: opt,
	}
}

func (c *simpleCore) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(c.com.Base, r)
}
func (c *simpleCore) Emit(w io.StreamWriter) error {
	c.com.Emitter.Bus(w)
	return c.com.Emitter.Work(w)
}

func (c *simpleCore) Apply(w io.StreamWriter, cg *types.Change) {
	bus, _ := c.com.Emitter.Bus(w)
	c.com.Base.Apply(bus, cg)
}
func (c *simpleCore) Delete(w io.StreamWriter) {
	bus, _ := c.com.Emitter.Bus(w)
	c.com.Base.Delete(bus)
	c.com.Emitter.Stop(w)
}

func (c *simpleCore) Command(ctx context.Context, args []string) (<-chan string, error) {
	return c.command(ctx, args)
}
