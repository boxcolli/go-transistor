package basicbase

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

// structs
type indexNode struct {
	Emitters map[emitter.Emitter]uint
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
	b.i = &indexNode{
		Emitters: make(map[emitter.Emitter]uint),
		Childs:   make(map[string]*indexNode),
	}
	b.icopy = &indexNode{
		Emitters: make(map[emitter.Emitter]uint),
		Childs:   make(map[string]*indexNode),
	}
	b.changes = make(chan *ecg)

	return b
}

// Base function implements
func (b *basicBase) Start() {
	stop := false
	for !stop {
		select {
		case cg := <-b.changes:
			// update before swap
			switch cg.Cg.Op {
			case types.OperationAdd:
				b.changeAdd(cg.Emitter, cg.Cg.Topics)
			case types.OperationDel:
				b.changeDel(cg.Emitter, cg.Cg.Topics)
			}

			// swap
			b.imx.Lock()
			b.i, b.icopy = b.icopy, b.i
			b.imx.Unlock()

			// update after swap
			switch cg.Cg.Op {
			case types.OperationAdd:
				b.changeAdd(cg.Emitter, cg.Cg.Topics)
			case types.OperationDel:
				b.changeDel(cg.Emitter, cg.Cg.Topics)
			}
		}
	}
}

func (b *basicBase) Stop() {
	panic("unimplemented")
}

func (b *basicBase) Flow(m *types.Message) error {
	topic := m.Topic

	// nil check
	if topic == nil || topic.Empty() {
		return base.ErrNoTopic
	}

	b.imx.Lock()
	curr := b.i
	exist := true
	for _, seg := range topic {
		child, ok := curr.Childs[seg]
		if !ok {
			exist = false
			break
		}
		curr = child
	}
	if exist {
		for e, _ := range curr.Emitters {
			e.Emit(m)
		}
	}
	b.imx.Unlock()

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
func (b *basicBase) changeAdd(e emitter.Emitter, topics []types.Topic) {
	// nil check
	if topics == nil || len(topics) == 0 {
		return
	}

	// process
	for _, topic := range topics {
		if topic == nil || topic.Empty() {
			continue
		}

		curr := b.icopy
		for _, seg := range topic {
			child, ok := curr.Childs[seg]
			if !ok {
				child = &indexNode{
					Emitters: make(map[emitter.Emitter]uint),
					Childs:   make(map[string]*indexNode),
				}
				child.Emitters[e] = 1
				curr.Childs[seg] = child
			} else if _, ok := child.Emitters[e]; !ok {
				child.Emitters[e] = 1
			} else {
				child.Emitters[e]++
			}
			curr = child
		}
	}
}

func (b *basicBase) changeDel(e emitter.Emitter, topics []types.Topic) {
	// nil check
	if topics == nil || len(topics) == 0 {
		return
	}

	// process
	for _, topic := range topics {
		if topic == nil {
			continue
		}

		curr := b.i
		exist := true
		for _, seg := range topic {
			// move to child
			child, ok := curr.Childs[seg]
			if !ok {
				exist = false
				break
			}
			curr = child

			// check is emitter in Emittes
			_, ex := curr.Emitters[e]
			if !ex {
				exist = false
				break
			}
		}
		if exist {
			curr.recurDel(e)
		}
	}
}

func (node *indexNode) recurDel(e emitter.Emitter) {
}
