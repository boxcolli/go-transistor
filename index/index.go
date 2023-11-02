package index

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type Index interface {
	Flow(m *types.Message)
	Add(e emitter.Emitter, t types.Topic) bool	// returns validity
	Del(e emitter.Emitter, t types.Topic) bool	// returns validity
}

type Inode struct {
	P		*Inode	// parent node
	Eset	map[emitter.Emitter]bool
	Next 	map[string]*Inode
}
func NewInode(P *Inode) *Inode {
	return &Inode{
		P,
		make(map[emitter.Emitter]bool),
		make(map[string]*Inode),
	}
}
func (n Inode) Empty() bool {
	return len(n.Eset) == 0 && len(n.Next) == 0
}
func (n *Inode) Emit(m *types.Message) {
	for e := range n.Eset {
		e.Emit(m)
	}
}

type Vnode struct {
	P		*Vnode
	Pair	*Inode
	Next	map[string]*Vnode
}
func NewVnode(P *Vnode, pair *Inode) *Vnode {
	return &Vnode{
		P,
		pair,
		make(map[string]*Vnode),
	}
}
func (n Vnode) Empty() bool {
	return len(n.Next) == 0
}
