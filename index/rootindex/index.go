package rootindex

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

// These methods are absolutely not goroutine safe.
// The caller must implement RW mutex between Flow and Swap.
// Calling RLock() while Flow() and Lock() while Swap() will be fine.
// Also, don't execute Apply(), Add() and Del() concurrently. Set one goroutine to control all.
type rootIndex struct {
	i	*index.Inode
	v 	map[emitter.Emitter]*index.Vnode
}

func NewRootIndex() index.Index {
	return &rootIndex{
		i: index.NewInode(),
		v: make(map[emitter.Emitter]*index.Vnode),
	}
}

// Flow implements Index.
func (I *rootIndex) Flow(m *types.Message) {
	i := I.i
	i.Emit(m)

	for _, token := range m.Topic {
		// Go to next
		if next, ok := i.Next[token]; !ok {
			break
		} else {
			i = next
		}

		i.Emit(m)
	}
}

// Add implements Index.
func (I *rootIndex) Add(e emitter.Emitter, t types.Topic) bool {
	var (
		i *index.Inode = I.i	// pointer to index
		v *index.Vnode		// pointer to inverted index
		x int = 0		// index of topic, indicating the first non-existing edge
	)

	if _, ok := I.v[e]; !ok {

		// Create headnode
		v = index.NewVnode(I.i)
		I.v[e] = v

		if t.Empty() {
			I.i.Eset[e] = true
			return true
		}

	} else {
		// Move v to the last existing node following the topic path
		v = I.v[e]
		branch := false	// Indicating that there is another branch in the path 
		for _, token := range t {
			next, ok := v.Next[token]
	
			if ok {
				if len(v.Next) > 1 {
					// Found another branch in the path
					branch = true
				}
				v = next
				x++
				continue
			}
	
			if v.Empty() && !branch {
				// The emitter is already listening to the supertopic
				// and trying to add a subtopic. (or add the current position)
				return false
			}
	
			i = v.Pair
			break
		}
	}

	if x == len(t) {
		// Delete subtree

		teardown(v, e, true)	// DFS, find leaves and erase (but don't erase the v)

		return true
	}

	{
		// Append subtree
		for ; x < len(t); x++ {
			token := t[x]
			next, ok := i.Next[token]

			if !ok {
				next = index.NewInode()
				i.Next[token] = next
				
				vnext := index.NewVnode(next)
				v.Next[token] = vnext

				i, v = next, vnext
				continue
			}

			vnext := index.NewVnode(next)
			v.Next[token] = vnext
			i, v = next, vnext
		}
		i.Eset[e] = true
		return true
	}
}

// Erase subtree of v
func teardown(v *index.Vnode, e emitter.Emitter, clean bool) {
	_teardown(v, e)
	if clean {
		v.Next = make(map[string]*index.Vnode)
	}
}
func _teardown(v *index.Vnode, e emitter.Emitter) {
	i := v.Pair

	// Touch leaf node and erase e from its index.Inode pair
	for token, vnext := range v.Next {
		if vnext.Empty() {
			// vnext is a leaf
			// delete e from inode pair
			inext := vnext.Pair
			delete(inext.Eset, e)

			if inext.Empty() {
				// Safe to delete the index.Inode
				delete(i.Next, token)
			}
		} else {
			_teardown(vnext, e)
		}
	}

	
}

// Del implements Index.
func (I *rootIndex) Del(e emitter.Emitter, t types.Topic) bool {
	var (
		v *index.Vnode		// pointer to vnode; the second last node
		bi *index.Inode		// the last inode that has another branch other than the topic path
		// bv *index.Vnode		// the last vnode that has another branch other than the topic path
	)
	if head, ok := I.v[e]; !ok {
		// The entry is empty
		return false
	} else {
		v = head
		bi = I.i
	}
	
	if t.Empty() {
		// Delete the entire entry
		teardown(v, e, false)
		delete(I.v, e)
		delete(I.i.Eset, e)
		return true
	}

	// Move v to the second last node
	x := 0
	for ; x < len(t) - 1; x++ {
		next, ok := v.Next[t[x]]

		if !ok {
			// The destination doesn't exist.
			return false
		}

		v = next

		// if len(next.Pair.Next) > 1 {
		// 	bi = next.Pair	// the next inode has another branch
		// }
		// if len(next.Next) > 1 {
		// 	bv = next		// the next vnode has another branch
		// }
	}

	{
		token := t[x]
		// v: the second last node
		// vnext: the last node
		if vnext, ok := v.Next[token]; !ok {
			// The destinatioin doesn't exist.
			return false
		} else {
			teardown(vnext, e, false)	// erase subtree of the last node
			delete(v.Next, token)		// erase the last node
			delete(vnext.Pair.Eset, e)	// 

			delete(bi.Next, token)
		}
	}

	return true
}
