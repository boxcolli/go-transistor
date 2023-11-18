package middleware

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type BaseMiddleware interface {
	Add(b base.Base)
	Del(b base.Base)
	base.Base
}

type baseMiddleware struct {
	base map[base.Base]bool
	mx sync.RWMutex
}
func NewBaseMiddleware(b base.Base) BaseMiddleware {
	mw := &baseMiddleware{
		base: make(map[base.Base]bool),
	}
	if b != nil {
		mw.Add(b)
	}
	return mw
}

func (mw *baseMiddleware) Add(b base.Base) {
	mw.mx.Lock()
	defer mw.mx.Unlock()
	mw.base[b] = true
}
func (mw *baseMiddleware) Del(b base.Base) {
	mw.mx.Lock()
	defer mw.mx.Unlock()
	delete(mw.base, b)
}

func (mw *baseMiddleware) Apply(e index.Entry, cg *types.Change) {
	mw.mx.RLock()
	defer mw.mx.RUnlock()
	for base := range mw.base {
		base.Apply(e, cg)
	}
}
func (mw *baseMiddleware) Delete(e index.Entry) {
	mw.mx.RLock()
	defer mw.mx.RUnlock()
	for base := range mw.base {
		base.Delete(e)
	}
}
func (mw *baseMiddleware) Flow(m *types.Message) {
	mw.mx.RLock()
	defer mw.mx.RUnlock()
	for base := range mw.base {
		base.Flow(m)
	}
}
