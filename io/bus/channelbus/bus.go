package channelbus

import (
	"sync"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type channelBus struct {
	ch chan *types.Message
	mx sync.Mutex
}

func NewChannelBus(qs int) io.Bus {
	return &channelBus{
		ch: make(chan *types.Message, qs),
		mx: sync.Mutex{},
	}
}
func (b *channelBus) Push(m *types.Message) {
	if ok := b.mx.TryLock(); !ok { return }	// this bus is turned off
	defer b.mx.Unlock()
	b.ch <- m
}
func (b *channelBus) Pull() <-chan *types.Message {
	return b.ch
}
func (b *channelBus) Lock() {
	b.mx.Lock()
}
func (b *channelBus) Unlock() {
	b.mx.Unlock()
}
