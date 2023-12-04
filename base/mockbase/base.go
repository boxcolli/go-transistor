package mockbase

import (
	"math/rand"
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type mockBase struct {
	e		map[int]index.Entry
	emx		sync.RWMutex
	inv		map[index.Entry]int
	invmx	sync.Mutex
}

func (b *mockBase) Apply(e index.Entry, cg *types.Change) {
	var me int
	{
		b.emx.Lock()
		for {
			me = rand.Int()
			if _, ok := b.e[me]; !ok {
				b.e[me] = e
				break
			}
		}
		b.emx.Unlock()
	}

	{
		b.invmx.Lock()
		b.inv[e] = me
		b.invmx.Unlock()
	}
}

func (b *mockBase) Delete(e index.Entry) {
	var me int
	{
		var ok bool
		b.invmx.Lock()
		me, ok = b.inv[e]
		if !ok {
			return
		}
		delete(b.inv, e)
		b.invmx.Unlock()
	}

	{
		b.emx.Lock()
		delete(b.e, me)
		b.emx.Unlock()
	}
}

func (b *mockBase) Flow(m *types.Message) {
	b.emx.RLock()
	defer b.emx.RUnlock()

	for _, e := range b.e {
		e.Push(m)
	}
}

func (b *mockBase) Start() {
	//
}

func (b *mockBase) Stop() {
	//
}

func NewMockBase() base.Base {
	return &mockBase{}
}
