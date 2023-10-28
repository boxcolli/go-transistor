package basiccore

import (
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/conductor"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/dialer"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/plug"
	"github.com/boxcolli/go-transistor/server"
	"github.com/boxcolli/go-transistor/types"
)

type BasicCore interface {
	core.Core
}

type Component struct {
	Collector collector.Collector
	Base      base.Base
	Emitter   emitter.Emitter

	Dialer    dialer.Dialer
	Plug      plug.Plug
	Conductor conductor.Conductor
	Server    server.Server
}

type CoreOption struct {
	Me     types.Member // Information about myself
	Cnames []string     // cluster names to dial at

}

type basicCore struct {
	com Component
	opt CoreOption
}

func NewCore(com Component, opt CoreOption) BasicCore {
	return &basicCore{}
}

// Start implements BasicCore.
func (*basicCore) Start() {
	panic("unimplemented")
}
