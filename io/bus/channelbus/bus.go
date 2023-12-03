package channelbus

import (
	"sync"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type channelBus struct {
	ch chan *types.Message
	ok bool	// status availability
	mx sync.RWMutex
}

func NewChannelBus(qs int) io.Bus {
	return &channelBus{
		ch: make(chan *types.Message, qs),
		mx: sync.RWMutex{},
	}
}
func (b *channelBus) Push(m *types.Message) {
	b.mx.RLock()
	defer b.mx.RUnlock()
	if !b.ok { return }	// this bus is turned off
	b.ch <- m
}
func (b *channelBus) Pull() <-chan *types.Message {
	return b.ch
}
func (b *channelBus) Lock() {
	b.mx.Lock()
	defer b.mx.Unlock()
	b.ok = false
}
func (b *channelBus) Unlock() {
	b.mx.Lock()
	defer b.mx.Unlock()
	b.ok = true
}
