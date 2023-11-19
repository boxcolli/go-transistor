package routeindex

import (
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type routeIndex struct {
	i	*index.Inode
	v	map[index.Entry](*index.Vnode)
}

// Flow implements index.Index.
func (I *routeIndex) Flow(m *types.Message) {
	i := I.i
	for _, seg := range m.Topic {
		if next, ok := i.Next[seg]; !ok {
			return
		} else {
			i = next
		}
	}

	i.Emit(m)
}

func NewRouteIndex() index.Index {
	return &routeIndex{
		i:	index.NewInode(nil),
		v:	make(map[index.Entry]*index.Vnode),
	}
}

// Add implements index.Index.
func (I *routeIndex) Add(e index.Entry, t types.Topic) bool {
	var (
		i *index.Inode = I.i
		v *index.Vnode
	)
	if vhead, ok := I.v[e]; !ok {
		vhead = index.NewVnode(nil, i)
		I.v[e] = vhead
		i.Eset[e] = true
		v = vhead
	} else {
		v = vhead
	}

	iappend := false
	vappend := false
	for _, seg := range t {
		if iappend {
			// Append and move both i & v			
			i = appendAndGetInext(i, e, seg)
			v = appendAndGetVnext(v, i, seg)
			continue
		}

		inext, ok := i.Next[seg]
		if !ok {
			// Append and move both i & v
			iappend = true
			i = appendAndGetInext(i, e, seg)
			v = appendAndGetVnext(v, i, seg)
			continue
		}
		i = inext	// Move i
		i.Eset[e] = true

		if vappend {
			// Append and move v
			v = appendAndGetVnext(v, i, seg)
			continue
		}
	
		vnext, ok := v.Next[seg]
		if !ok {
			// Append and move v
			vappend = true
			v = appendAndGetVnext(v, i, seg)
			continue
		}
		v = vnext	// Move v
	}

	return iappend || vappend
}

func appendAndGetInext(i *index.Inode, e index.Entry, t string) *index.Inode {
	// i.Eset[e] = true
	inext := index.NewInode(i)
	inext.Eset[e] = true
	i.Next[t] = inext
	return inext
}

func appendAndGetVnext(v *index.Vnode, inext *index.Inode, t string) *index.Vnode {
	vnext := index.NewVnode(v, inext)
	v.Next[t] = vnext
	return vnext
}

// Del implements index.Index.
func (I *routeIndex) Del(e index.Entry, topic types.Topic) bool {
	var (
		v *index.Vnode
	)
	{
		if vhead, ok := I.v[e]; !ok {
			// Nothing to delete
			return false
		} else {
			v = vhead
		}
	}

	if topic.Empty() {
		// Wipe out e from entire index
		teardown(I.v[e], e)
		delete(I.v, e)
		return true
	}

	// Move v to the last node along the topic path
	for x := 0; x < len(topic); x++ {
		seg := topic[x]
		vnext, ok := v.Next[seg]
		if !ok {
			return false
		}

		v = vnext
	}

	if !v.Empty() {
		// The path should not be deleted
		return false
	}

	vlast := tearup(v, e, topic, len(topic) - 1)
	if vlast == I.v[e] {
		delete(I.v, e)
	}

	return true
}

func tearup(v *index.Vnode, e index.Entry, topic types.Topic, x int) *index.Vnode {
	pv := v.Prev
	if pv == nil {
		delete(v.Pair.Eset, e)
		return v
	}

	delete(v.Pair.Eset, e)		// delete e from current inode
	delete(pv.Next, topic[x])	// delete current vnode
	if v.Pair.Empty() {
		delete(pv.Pair.Next, topic[x])	// delete current inode from parent inode
	}

	if pv.Empty() {
		return tearup(pv, e, topic, x - 1)
	}
	return nil
}

// Erase subtree of v
func teardown(v *index.Vnode, e index.Entry) {
	_teardown(v, e, "")
}

func _teardown(v *index.Vnode, e index.Entry, myKey string) {
	if !v.Empty() {
		for key, vnext := range v.Next {
			_teardown(vnext, e, key)
		}
	}

	pv := v.Prev
	if pv == nil {
		delete(v.Pair.Eset, e)
		return
	}

	delete(v.Pair.Eset, e)	// delete e from current inode
	delete(pv.Next, myKey)	// delete current vnode
	if v.Pair.Empty() {
		delete(pv.Pair.Next, myKey)	// delete current inode from parent inode
	}
}
