package basiccore

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/conductor"
	"github.com/boxcolli/go-transistor/core"
	"github.com/boxcolli/go-transistor/dialer"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/plug"
	"github.com/boxcolli/go-transistor/types"
)

type BasicCore interface {
	core.Core
}

type CoreComponent struct {
	Collector	*collector.Collector
	Base		*base.Base
	Emitter		*emitter.Emitter

	Dialer		*dialer.Dialer
	Plug 		*plug.Plug
	Conductor	*conductor.Conductor
	Server		*pb.TransistorServiceServer
}

type CoreOption struct {
	Me types.Member	// Information about myself
	Cnames []string	// cluster names to dial at

}

type basicCore struct {
	com		CoreComponent
	opt		CoreOption
}

func NewCore(com CoreComponent, opt CoreOption) BasicCore {
	return &basicCore{}
}
