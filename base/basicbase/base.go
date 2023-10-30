package basicbase

// import (
// 	"sync"

// 	"github.com/boxcolli/go-transistor/base"
// 	"github.com/boxcolli/go-transistor/emitter"
// 	"github.com/boxcolli/go-transistor/types"
// )

// type ecg struct {
// 	e	emitter.Emitter
// 	cg  *types.Change
// }

// type basicBase struct {
// 	i       *inode	// index {Key: topic, Value: set(Emitter)}
// 	icopy   *inode
// 	inv		map[emitter.Emitter]*vnode // inverted {Key: Emitter, Value: set(topic)}
// 	imx     sync.RWMutex

// 	changes chan *ecg
// }

// // Create BasicBase instance
// func NewBasicBase(qsiz int) base.Base {
// 	return &basicBase{
// 		i:			newInode(),
// 		icopy:		newInode(),
// 		inv:		make(map[emitter.Emitter]*vnode),
// 		imx:		sync.RWMutex{},
// 		changes:	make(chan *ecg, qsiz),
// 	}
// }

// // Base function implements
// func (b *basicBase) Start() {
// 	stop := false
// 	for !stop {
// 		select {
// 		case cg := <-b.changes:
// 			// update before swap
// 			switch cg.cg.Op {
// 			case types.OperationAdd:
// 				b.applyAdd(cg.e, cg.cg.Topics)
// 			case types.OperationDel:
// 				b.applyDel(cg.e, cg.cg.Topics)
// 			}

// 			// swap
// 			b.imx.Lock()
// 			b.i, b.icopy = b.icopy, b.i
// 			b.imx.Unlock()

// 			// update after swap
// 			switch cg.cg.Op {
// 			case types.OperationAdd:
// 				b.applyAdd(cg.e, cg.cg.Topics)
// 			case types.OperationDel:
// 				b.applyDel(cg.e, cg.cg.Topics)
// 			}
// 		}
// 	}
// }

// func (b *basicBase) Stop() {
// 	panic("unimplemented")
// }

// func (b *basicBase) Flow(m *types.Message) error {
// 	b.imx.RLock()
// 	defer b.imx.RUnlock()

// 	curr := b.i
// 	{
// 		// Flow message
// 		for e := range curr.eset {
// 			e.Emit(m)
// 		}
	
// 		// Traverse tree
// 		for _, seg := range m.Topic {
// 			// Go to next node
// 			var ok bool
// 			if curr, ok = curr.next[seg]; !ok {
// 				break
// 			}
	
// 			// Flow message
// 			for e := range curr.eset {
// 				e.Emit(m)
// 			}
// 		}
// 	}

// 	return nil
// }

// func (b *basicBase) Apply(e emitter.Emitter, cg *types.Change) {
// 	b.changes <- &ecg{ e, cg }
// }

// func (b *basicBase) Delete(e emitter.Emitter) {
// 	b.Apply(e, &types.Change{
// 		Op: types.OperationDel,
// 		Topics: []types.Topic{{}},
// 	})
// }

// // basicbase functions
// func (b *basicBase) applyAdd(e emitter.Emitter, topics []types.Topic) {
// 	// Find difference in inverted index
// 	diff := b.inv[e]
	
	


// 	for _, topic := range topics {
// 		curr := b.icopy
// 		for _, seg := range topic {
// 			next, ok := curr.next[seg]
// 			if !ok {
// 				// Append new child
// 				curr.next[seg] = &inode{
// 					eset: make(map[emitter.Emitter]bool),
// 					next:   make(map[string]*inode),
// 				}
// 			}
// 			next.eset[e] = true
// 			curr = next
// 		}
// 	}
// }

// func (b *basicBase) applyDel(e emitter.Emitter, topics []types.Topic) {
// 	// nil check
// 	if len(topics) == 0 {
// 		return
// 	}

// 	// process
// 	for _, topic := range topics {
// 		if topic.Empty() {
// 			b.icopy.recurDel(e, "", nil)
// 			continue
// 		}

// 		curr := b.i
// 		var parent *inode = nil

// 		exist := true
// 		for _, seg := range topic {
// 			// move to child
// 			child, ok := curr.next[seg]
// 			if !ok {
// 				exist = false
// 				break
// 			}
// 			parent = curr
// 			curr = child

// 			// check if emitter is in eset
// 			_, ex := curr.eset[e]
// 			if !ex {
// 				exist = false
// 				break
// 			}
// 		}
// 		if exist {
// 			curr.recurDel(e, topic[len(topic)-1], parent)
// 		}
// 	}
// }

// func (node *inode) recurDel(e emitter.Emitter, key string, parent *inode) {
// 	if _, ex := node.eset[e]; parent == nil || ex {
// 		for key, child := range node.next {
// 			child.recurDel(e, key, node)
// 		}
// 		if parent != nil {
// 			delete(node.eset, e)
// 			if len(node.eset) == 0 {
// 				delete(parent.next, key)
// 			}
// 		}
// 	}
// }
