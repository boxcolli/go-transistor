package rootindex2

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type rootIndex struct {
	i *index.Inode                     // index
	v map[emitter.Emitter]*index.Vnode // inverted index
}

func NewRootIndex() index.Index {
	return &rootIndex{
		i: index.NewInode(),
		v: make(map[emitter.Emitter]*index.Vnode),
	}
}

// Find the destination node, emit throughout the node and its subtree nodes
func (I *rootIndex) Flow(m *types.Message) {
	i := I.i
	for _, t := range m.Topic {
		if inext, ok := i.Next[t]; !ok {
			return
		} else {
			i = inext
		}
	}
	flow(i, m)
}

// Recursively emit m throughout the i and its subtree nodes
func flow(i *index.Inode, m *types.Message) {
	i.Emit(m)
	for _, inext := range i.Next {
		flow(inext, m)
	}
}

func (I *rootIndex) Add(e emitter.Emitter, topic types.Topic) bool {
	var v *index.Vnode
	var x int = 0

	if vhead, ok := I.v[e]; !ok {
		// e is added for first time
		v = index.NewVnode(I.i)
		I.v[e] = v
	} else {
		v = vhead
		if v.Empty() {
			// e is already listening to the entire stream.
			return false
		}

		// Move v to the last existing node following the topic path
		for ; x < len(topic); x++ {
			if vnext, ok := v.Next[topic[x]]; !ok {
				break
			} else {
				v = vnext
			}
		}
		
		if v.Empty() {
			// e is already listening to the supertopic
			return false
		}

		if x == len(topic) {
			// e is adding a supertopic
			teardown(v, e, true)
			return true
		}
	}

	// Append
	var i = I.i
	var iappend = false
	for ; x < len(topic); x++ {
		t := topic[x]

		if iappend {
			i = appendAndGetInext(i, e, t)
			v = appendAndGetVnext(v, i, t)
			continue
		}

		// Move i
		inext, ok := i.Next[t]
		if !ok {
			iappend = true
			i = appendAndGetInext(i, e, t)
			v = appendAndGetVnext(v, i, t)
			continue
		}
		i = inext

		// Move v
		v = appendAndGetVnext(v, i, t)
	}

	// Add e
	i.Eset[e] = true

	return true
}

func appendAndGetInext(i *index.Inode, e emitter.Emitter, t string) *index.Inode {
	inext := index.NewInode()
	i.Next[t] = inext
	return inext
}

func appendAndGetVnext(v *index.Vnode, inext *index.Inode, t string) *index.Vnode {
	vnext := index.NewVnode(inext)
	v.Next[t] = vnext
	return vnext
}

// Erase subtree of v (but not v)
func teardown(v *index.Vnode, e emitter.Emitter, cleanV bool) {
	_teardown(v, e)
	if cleanV {
		v.Next = make(map[string]*index.Vnode)
	}
}
func _teardown(v *index.Vnode, e emitter.Emitter) {
	i := v.Pair

	for t, vnext := range v.Next {
		if vnext.Empty() {

			// vnext is a leaf; delete e from inode pair
			inext := vnext.Pair
			delete(inext.Eset, e)

			if inext.Empty() {
				// Safe to delete the index.Inode
				delete(i.Next, t)
			}

		} else {
			_teardown(vnext, e)
		}
	}
}

func (I *rootIndex) Del(e emitter.Emitter, topic types.Topic) bool {
	var v *index.Vnode
	var bi *index.Inode
	var bix int
	var bv *index.Vnode

	if vhead, ok := I.v[e]; !ok {
		return false
	} else {
		v = vhead
	}

	if topic.Empty() {
		teardown(v, e, false)
		delete(I.v, e)
		delete(I.i.Eset, e)
		return true
	}

	// Move v to the last node
	{
		x := 0
		if len(v.Pair.Next) > 1 || len(v.Pair.Eset) > 0 {
			bi = v.Pair	// v.Pair has another branch
			bix = x
		}
		if len(v.Next) > 1 {
			bv = v		// v has another branch
		}
		for ; x < len(topic); x++ {
			vnext, ok := v.Next[topic[x]]
			if !ok {
				// The path doesn't exist.
				return false
			}
			v = vnext
	
			if len(v.Pair.Next) > 1 || len(v.Pair.Eset) > 0 {
				bi = v.Pair	// v.Pair has another branch
				bix = x
			}
			if len(v.Next) > 1 {
				bv = v		// v has another branch
			}
		}
	}

	if bv != nil {
		// Delete bv ~
		teardown(bv, e, true)
	} else {
		// Deleve v ~
		teardown(v, e, true)
	}

	// Delete bi ~
	if bi != nil {
		delete(bi.Next, topic[bix])
	}

	return true
}
