package basicindex

import (
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type basicIndex struct {
	i *index.Inode
	v map[index.Entry](*index.Vnode)
}

func (*basicIndex) Del(e index.Entry, t types.Topic) bool {
	panic("unimplemented")
}

func NewBasicIndex() index.Index {
	return &basicIndex{
		i: index.NewInode(nil),
		v: make(map[index.Entry]*index.Vnode),
	}
}

func (I *basicIndex) Flow(m *types.Message) {
	i := I.i
	i.Emit(m)

	// Go to the destination
	for _, seg := range m.Topic {
		if next, ok := i.Next[seg]; !ok {
			return
		} else {
			i = next
		}

		i.Emit(m)
	}

	flow(i, m)
}

// Recursively emit
func flow(i *index.Inode, m *types.Message) {
	for _, inext := range i.Next {
		inext.Emit(m)
		flow(inext, m)
	}
}

func (I *basicIndex) Add(e index.Entry, t types.Topic) bool {
	var (
		i *index.Inode = I.i
		v *index.Vnode
	)
	{
		if vhead, ok := I.v[e]; !ok {
			v = index.NewVnode(nil, I.i)
			I.v[e] = v
		} else {
			// if vhead.Empty() { return false }
			v = vhead
		}
	}

	// Move v
	var (
		iappend = false
		vappend = false
		branch = false
	)
	if v.Empty() { }
	for _, seg := range t {
		if iappend {
			i = appendAndGetInext(i, e, seg)
			v = appendAndGetVnext(v, i, seg)
			continue
		}

		inext, ok := i.Next[seg]
		if !ok {

		}
	}
}

func appendAndGetInext(i *index.Inode, e index.Entry, t string) *index.Inode {
	inext := index.NewInode(i)
	i.Next[t] = inext
	return inext
}

func appendAndGetVnext(v *index.Vnode, inext *index.Inode, t string) *index.Vnode {
	vnext := index.NewVnode(v, inext)
	v.Next[t] = vnext
	return vnext
}
