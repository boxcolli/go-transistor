package base

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

// These methods are absolutely not goroutine safe.
// The caller must implement RW mutex between Flow and Swap.
// Calling RLock() while Flow() and Lock() while Swap() will be fine.
// Also, don't execute Apply(), Add() and Del() concurrently. Set one goroutine to control all.
type Index interface {
	Flow(m *types.Message)
	Add(e emitter.Emitter, t types.Topic) bool	// returns validity
	Del(e emitter.Emitter, t types.Topic) bool	// returns validity
}

type inode struct {
	eset map[emitter.Emitter]bool
	next map[string]*inode
}
func newInode() *inode {
	return &inode{
		make(map[emitter.Emitter]bool),
		make(map[string]*inode),
	}
}
func (n inode) empty() bool {
	return len(n.eset) == 0 && len(n.next) == 0
}
func (n inode) emit(m *types.Message) {
	for e, _ := range n.eset {
		e.Emit(m)
	}
}

type vnode struct {
	pair *inode
	next map[string]*vnode
}
func newVnode(pair *inode) *vnode {
	return &vnode{
		pair,
		make(map[string]*vnode),
	}
}
func (n vnode) empty() bool {
	return len(n.next) == 0
}

type basicIndex struct {
	i	*inode                     	// index
	v 	map[emitter.Emitter]*vnode	// inverted index
}

func NewBasicIndex() Index {
	return &basicIndex{
		i: newInode(),
		v: make(map[emitter.Emitter]*vnode),
	}
}

// Flow implements Index.
func (i *basicIndex) Flow(m *types.Message) {
	curr := i.i
	curr.emit(m)

	for _, token := range m.Topic {
		// Go to next
		if next, ok := curr.next[token]; !ok {
			break
		} else {
			curr = next
		}

		// Send message
		curr.emit(m)
	}

	// Send message
	curr.emit(m)
}

// Add implements Index.
func (I *basicIndex) Add(e emitter.Emitter, t types.Topic) bool {
	// Initialize
	var (
		i *inode = I.i	// pointer to index
		v *vnode		// pointer to inverted index
		x int = 0		// index of topic, indicating the first non-existing edge
	)

	if _, ok := I.v[e]; !ok {
		// Create headnode
		I.v[e] = newVnode(I.i)
		I.i.eset[e] = true
	} else {
		// Move v to the last existing node following the topic path
		v = I.v[e]
		branch := false	// If branch is not found, the 
		for _, token := range t {
			next, ok := v.next[token]
	
			if ok {
				if len(v.next) > 1 {
					// Found another branch in the path
					branch = true
				}
				v = next
				x++
				continue
			}
	
			if v.empty() && !branch {
				// The emitter is already listening to the supertopic
				// and trying to add a subtopic. (or add the current position)
				return false
			}
	
			i = v.pair
			break
		}
	}

	if x == len(t) {
		// Delete subtree

		teardown(e, v)	// DFS, find leaves and erase (but don't erase the v)

		return true
	}

	{
		// Append subtree
		for ; x < len(t); x++ {
			token := t[x]
			next, ok := i.next[token]

			if !ok {
				next = newInode()
				i.next[token] = next
				
				vnext := newVnode(next)
				v.next[token] = vnext

				i, v = next, vnext
				continue
			}

			vnext := newVnode(next)
			v.next[token] = vnext
			i, v = next, vnext
		}
		i.eset[e] = true
		return true
	}
}

func teardown(e emitter.Emitter, v *vnode) {
	if v.empty() {
		i := v.pair
		delete(i.eset, e)
		if i.empty() {
			
		}
		return
	}

	for token, next := range v.next {
		if next.empty() {
			i := next.pair
			delete(i.eset, e)
			if i.empty() {
				delete(v.pair.next, token)
			}
		}
	}
}

// Del implements Index.
func (i *basicIndex) Del(e emitter.Emitter, t types.Topic) bool {
	panic("unimplemented")
}
