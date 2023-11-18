package spoofmiddleware

import (
	"sync"

	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/types"
)

type SpoofMiddleware interface {
	Stop()
	base.Base
}

type spoofMiddleware struct {
	ch	chan<- *types.Message
	mx	sync.Mutex
}

func NewCountMiddlware(ch chan<- *types.Message) SpoofMiddleware {
	return &spoofMiddleware{
		ch: ch,
	}
}
func (mw *spoofMiddleware) Stop() {
	mw.mx.Lock()
}
func (*spoofMiddleware) Apply(e index.Entry, cg *types.Change) {}
func (*spoofMiddleware) Delete(e index.Entry) {}
func (mw *spoofMiddleware) Flow(m *types.Message) {
	if ok := mw.mx.TryLock(); !ok { return }
	mw.ch <- m
	mw.mx.Unlock()
}
