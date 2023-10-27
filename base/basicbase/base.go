package basicbase

import (
	"fmt"
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type indexNode struct {
	Emitters []emitter.Emitter
	Childs   map[string]*indexNode
}

type ecg struct {
	Emitter emitter.Emitter
	Cg      *types.Change
}

type basicBase struct {
	i       map[string]*indexNode
	icopy   map[string]*indexNode
	imx     sync.RWMutex
	changes chan *ecg
}

func NewBasicBase() base.Base {
	return &basicBase{}
}

func (b *basicBase) Start() {
	b.changes = make(chan *ecg)
	go b.changeLoop()
}

func (b *basicBase) Stop() {
	panic("unimplemented")
}

func (b *basicBase) changeLoop() {
	stop := false
	for !stop {
		select {
		case cg := <-b.changes:
			b.imx.RLock()
			switch cg.Cg.Op {
			case types.OperationAdd:
				fmt.Print("Add")
			case types.OperationDel:
				fmt.Print("Del")
			}
			b.imx.Unlock()
		}
	}
}

func (b *basicBase) Flow(m *types.Message) error {
	return nil
}

func (b *basicBase) Apply(e emitter.Emitter, cg *types.Change) {
	b.changes <- &ecg{
		e,
		cg,
	}
}

func (b *basicBase) Delete(e emitter.Emitter) {
	b.Apply(e, &types.Change{
		Op:     types.OperationDel,
		Topics: nil,
	})
}
