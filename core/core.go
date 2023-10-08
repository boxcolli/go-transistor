package core

import (
	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/dialer"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
)

// A core is a builder and also a CLI engine
type Core interface {

}

type CoreComponent struct {
	Collector	*collector.Collector
	Base		*base.Base
	Emitter		*emitter.Emitter
	Dialer		*dialer.Dialer
	Plug 		*plugs.Plug
	Server		*pb.TransistorServiceServer
}

type CoreOption struct {
	Me types.Member	// Information about myself
	Cnames []string	// cluster names to dial at

}

type core struct {
	com		CoreComponent
	opt		CoreOption
}

func NewCore(com CoreComponent, opt CoreOption) Core {
	return &core{}
}
