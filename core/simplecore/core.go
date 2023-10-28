package simplecore

import (
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/server"
	"github.com/boxcolli/go-transistor/types"
)

type Component struct {
	Collector collector.Collector
	Base      base.Base
	Emitter   emitter.Emitter

	Server    server.Server
}

type Option struct {
	Topics []types.Topic
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
