package basicbase

import (
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

// Node for index tree
type inode struct {
	pair *inode
	eset map[emitter.Emitter]bool
	next map[string]*inode
}

func newInodePair() (*inode, *inode) {
	a, b := newInode(), newInode()
	a.pair, b.pair = b, a
	return a, b
}

func newInode() *inode {
	return &inode{
		pair: nil,
		eset: make(map[emitter.Emitter]bool),
		next: make(map[string]*inode),
	}
}

func (n inode) emptry() bool { return len(n.eset) == 0 }
func (n inode) isLeaf() bool { return len(n.next) == 0 }

// Node for inverted index tree or single item trace tree
type vnode struct {
	entry	*inode
	next	map[string]*vnode
}

func newVnode(entry *inode) *vnode {
	return &vnode{
		entry: entry,
		next: make(map[string]*vnode),
	}
}

type structure struct {
	index	*inode
	copy	*inode
	inv		map[emitter.Emitter]*vnode
}

func (s structure) flow(m *types.Message) {

}

func (s *structure) add(e emitter.Emitter, t types.Topic) {
	
}

func (s *structure) del(e emitter.Emitter, t types.Topic) {

}

// This method is absolutely not safe.
// The caller must apply mutex mechanism.
func (s *structure) swap() {
	i.i, i.icopy = i.icopy, i.i
}

// Flush differences into icopy
func (s *structure) flush() {
	
}