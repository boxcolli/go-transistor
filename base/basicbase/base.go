package basicbase

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type indexNode struct {
	emitters []*emitter.Emitter
	childs   map[string]*indexNode
}

type basicBase struct {
	i       map[string]*indexNode
	icopy   map[string]*indexNode
	imx     sync.RWMutex
	changes chan *types.Change
}

func NewBasicBase() base.Base {
	return &basicBase{}
}

func (b *basicBase) Start() {
	b.changes = make(chan *types.Change)
	go b.changeLoop()
}

func (b *basicBase) Stop() {
	panic("unimplemented")
}

func (b *basicBase) changeLoop() {
}

func (b *basicBase) Flow(m *types.Message) error {
	return nil
}

func (b *basicBase) Apply(e *emitter.Emitter, cg *types.Change) {
	panic("unimplemented")
}

func (b *basicBase) Delete(e *emitter.Emitter) {
	panic("unimplemented")
}
