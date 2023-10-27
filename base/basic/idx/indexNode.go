package basicbaseidx

import "github.com/boxcolli/go-transistor/emitter"

type IndexNode struct {
	emitters []*emitter.Emitter
	childs   map[string]*IndexNode
}
