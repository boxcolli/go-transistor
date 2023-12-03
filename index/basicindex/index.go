package basicindex

import (
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type basicIndex struct {
	i *index.Inode
	v map[index.Entry]*index.Vnode
}

func NewBasicIndex() index.Index {
	return &basicIndex{
		i: index.NewInode(nil, ""),
		v: make(map[index.Entry]*index.Vnode),
	}
}

// Push message to every entry that is on the topic path.
func (I *basicIndex) Flow(m *types.Message) {
	i := I.i
	i.Push(m)

	for _, token := range m.Topic {
		if next, ok := i.Next[token]; !ok {
			break
		} else {
			i = next
		}

		i.Push(m)
	}
}

func (I *basicIndex) Add(e index.Entry, t types.Topic) bool {
	var (
		i *index.Inode = I.i
		v *index.Vnode
		iappend = false
		vappend = false
	)
	if vroot, ok := I.v[e]; !ok {
		vroot = index.NewVnode(nil, "", i)
		I.v[e] = vroot
		v = vroot
		vappend = true
	} else {
		v = vroot
	}

	// Move i & v to the destination
	for _, token := range t {
		if iappend {
			// Append and move both i & v			
			i = appendAndGetInext(i, e, token)
			v = appendAndGetVnext(v, i, token)
			continue
		}
		inext, ok := i.Next[token]
		if !ok {
			// Check branch.
			// 	If there's no other branch, the entry already subscribed the supertopic.
			// 	If the vnode is already appended, this condition is not necessary.
			if !vappend && len(v.Next) == 0 { return false }

			iappend = true
			vappend = true
			i = appendAndGetInext(i, e, token)
			v = appendAndGetVnext(v, i, token)
			continue
		}
		i = inext

		if vappend {
			// Append and move v
			v = appendAndGetVnext(v, i, token)
			continue
		}
		vnext, ok := v.Next[token]
		if !ok {
			// Check branch
			if len(v.Next) == 0 { return false }

			vappend = true
			v = appendAndGetVnext(v, i, token)
			continue
		}
		v = vnext
	}

	if len(v.Next) == 0 && !vappend {
		return false
	}

	// Mark entry at the current position
	i.Eset[e] = true

	I.deleteSubtree(e, v)

	return true
}
func appendAndGetInext(i *index.Inode, e index.Entry, t string) *index.Inode {
	inext := index.NewInode(i, t)
	i.Next[t] = inext
	return inext
}

func appendAndGetVnext(v *index.Vnode, inext *index.Inode, t string) *index.Vnode {
	vnext := index.NewVnode(v, t, inext)
	v.Next[t] = vnext
	return vnext
}

// Remove the destination node(d) and its subtree.
// If the parent node of d doesn't have other branch,
// erase the parent recursively until there is other branch.
func (I *basicIndex) Del(e index.Entry, t types.Topic) bool {
	var (
		v *index.Vnode
	)
	if vroot, ok := I.v[e]; !ok {
		return false	// There's nothing to delete
	} else {
		v = vroot
	}

	// Move v
	for _, token := range t {
		vnext, ok := v.Next[token]
		if !ok { return false }
		v = vnext
	}

	delete(v.Pair.Eset, e)
	I.deleteSubtree(e, v)
	deleteIBranch(v.Pair)
	I.deleteVBranch(e, v, nil)
	return true
}

func (I *basicIndex) deleteSubtree(e index.Entry, v *index.Vnode) {
	for _, vnext := range v.Next {
		I._deleteSubtree(e, vnext, v)
	}
}

// Go to leaf, delete e from i and delete branch
func (I *basicIndex) _deleteSubtree(e index.Entry, v *index.Vnode, stop *index.Vnode) {
	if v.Empty() {
		// v is leaf
		delete(v.Pair.Eset, e)
		deleteIBranch(v.Pair)
		I.deleteVBranch(e, v, stop)
		return
	}

	// Go to leaf
	for _, vnext := range v.Next {
		I.deleteSubtree(e, vnext)
	}
}

func deleteIBranch(i *index.Inode) {
	if i.Prev != nil && i.Empty() {
		delete(i.Prev.Next, i.Token)	// delete i from prev
		deleteIBranch(i.Prev)			// move to prev
	}
}

func (I *basicIndex) deleteVBranch(e index.Entry, v *index.Vnode, stop *index.Vnode) {
	if v == stop || !v.Empty() { return }

	// v is empty
	if v.Prev == nil {
		// if _, ok := v.Pair.Eset[e]; !ok {
		// }
		delete(I.v, e)	// delete vroot of e
	} else {
		delete(v.Prev.Next, v.Token)		// delete v from prev
		I.deleteVBranch(e, v.Prev, stop)	// move to prev
	}
}
