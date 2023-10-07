package core

import (
	"github.com/boxcolli/go-transistor/emitters"
	"github.com/boxcolli/go-transistor/plugs"
)

// A core is a builder and also a CLI engine
type Core interface {

}

type CoreOption struct {
	cname string
	name string
	
	p plugs.Plug
	e emitters.Emitter
}

type core struct {
	
}

func NewCore(opt CoreOption) {
	
}
