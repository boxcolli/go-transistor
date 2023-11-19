package countmiddleware

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type CountMiddleware interface {
	Stop()
	base.Base
}

type countMiddleware struct {
	ch	chan<- bool
	mx	sync.Mutex
}

func NewCountMiddlware(ch chan<- bool) CountMiddleware {
	return &countMiddleware{
		ch: ch,
	}
}
func (mw *countMiddleware) Stop() {
	mw.mx.Lock()
}
func (*countMiddleware) Apply(index.Entry, *types.Change) {}
func (*countMiddleware) Delete(index.Entry) {}
func (mw *countMiddleware) Flow(*types.Message) {
	if ok := mw.mx.TryLock(); !ok { return }
	mw.ch <- true
	mw.mx.Unlock()
}
