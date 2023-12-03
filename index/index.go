package index

import (
	"github.com/boxcolli/go-transistor/types"
)

type Entry interface {
	Push(m *types.Message)
}

type Index interface {
	Flow(m *types.Message)
	Add(e Entry, t types.Topic) bool	// returns validity
	Del(e Entry, t types.Topic) bool	// returns validity
}

type Inode struct {
	Prev	*Inode	// parent node
	Token	string
	Eset	map[Entry]bool
	Next 	map[string]*Inode
}
func NewInode(prev *Inode, token string) *Inode {
	return &Inode{
		prev,
		token,
		make(map[Entry]bool),
		make(map[string]*Inode),
	}
}
func (n Inode) Empty() bool {
	return len(n.Eset) == 0 && len(n.Next) == 0
}
func (n *Inode) Push(m *types.Message) {
	for e := range n.Eset {
		e.Push(m)
	}
}

type Vnode struct {
	Prev	*Vnode
	Token	string
	Pair	*Inode
	Next	map[string]*Vnode
}
func NewVnode(prev *Vnode, token string, pair *Inode) *Vnode {
	return &Vnode{
		prev,
		token,
		pair,
		make(map[string]*Vnode),
	}
}
func (n Vnode) Empty() bool {
	return len(n.Next) == 0
}
