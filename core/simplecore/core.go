package simplecore

import (
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type Component struct {
	Collector collector.Collector
	Base      base.Base
	Emitter   emitter.Emitter
}

type Option struct {
	StaticTopics []types.Topic
}

type simpleCore struct {
	com Component
	opt Option
}

func NewSimpleCore(com Component, opt Option) core.Core {
	return &simpleCore{
		com: com,
		opt: opt,
	}
}

// Start implements core.Core.
func (c *simpleCore) Start() {
	panic("unimplemented")
}

// Apply implements core.Core.
func (c *simpleCore) Apply(e emitter.Emitter, cg *types.Change) {
	c.com.Base.Apply(e, cg)
}

// Collect implements core.Core.
func (c *simpleCore) Collect(r io.StreamReader) error {
	return c.com.Collector.Work(r)
}

// Command implements core.Core.
func (c *simpleCore) Command(args []string) chan string {
	panic("unimplemented")
}

// Delete implements core.Core.
func (c *simpleCore) Delete(e emitter.Emitter) {
	c.com.Base.Delete(e)
}

// Emit implements core.Core.
func (c *simpleCore) Emit(emitter.Emitter, io.StreamWriter) error {
	
}
