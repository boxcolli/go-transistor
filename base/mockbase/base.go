package mockbase

import (
	"math/rand"
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/emitter"
	"github.com/boxcolli/go-transistor/types"
)

type mockBase struct {
	e		map[int]emitter.Emitter
	emx		sync.RWMutex
	inv		map[emitter.Emitter]int
	invmx	sync.Mutex
}

func (b *mockBase) Apply(e emitter.Emitter, cg *types.Change) {
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

func (b *mockBase) Delete(e emitter.Emitter) {
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
		e.Emit(m)
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
