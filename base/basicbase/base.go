package basicbase

import (
	"fmt"
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

// structs
type indexNode struct {
	Emitters map[emitter.Emitter]bool
	Childs   map[string]*indexNode
}

type ecg struct {
	Emitter emitter.Emitter
	Cg      *types.Change
}

type basicBase struct {
	i       *indexNode
	icopy   *indexNode
	imx     sync.RWMutex
	changes chan *ecg
}

// Create BasicBase instance
func NewBasicBase() base.Base {
	b := &basicBase{}
	b.i = &indexNode{}
	b.icopy = &indexNode{}
	b.changes = make(chan *ecg)

	return b
}

// Base function implements
func (b *basicBase) Start() {
	stop := false
	for !stop {
		select {
		case cg := <-b.changes:
			b.imx.RLock()

			// update before swap
			switch cg.Cg.Op {
			case types.OperationAdd:
				b.changeAdd(cg.Emitter, cg.Cg.Topics)
			case types.OperationDel:
				b.changeDel(cg.Emitter, cg.Cg.Topics)
			}

			// swap
			b.i, b.icopy = b.icopy, b.i

			// update after swap
			switch cg.Cg.Op {
			case types.OperationAdd:
				b.changeAdd(cg.Emitter, cg.Cg.Topics)
			case types.OperationDel:
				b.changeDel(cg.Emitter, cg.Cg.Topics)
			}

			b.imx.Unlock()
		}
	}
}

func (b *basicBase) Stop() {
	panic("unimplemented")
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

// basicbase functions
func (b *basicBase) changeAdd(emitter emitter.Emitter, topics []types.Topic) {
	fmt.Print("changeAdd()")
}

func (b *basicBase) changeDel(emitter emitter.Emitter, topics []types.Topic) {
	fmt.Print("changeDel()")
}
