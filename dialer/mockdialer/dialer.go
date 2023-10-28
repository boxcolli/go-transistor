package mockdialer

import (
	"context"
	"sync"

	"github.com/boxcolli/go-transistor/dialer"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/io/reader/channelreader"
	"github.com/boxcolli/go-transistor/types"
)

type mockDialer struct {
	m map[*types.Member]io.StreamReader
	mx sync.Mutex
	ch map[*types.Member]chan *types.Message
}

func NewMockDialer(m map[*types.Member]io.StreamReader, ch map[*types.Member]chan *types.Message) dialer.Dialer {
	return &mockDialer{
		m: m,
		ch: ch,
		mx: sync.Mutex{},
	}
}

// Dial implements dialer.Dialer.
func (d *mockDialer) Dial(ctx context.Context, m *types.Member) (io.StreamReader, error) {
	d.mx.Lock()
	defer d.mx.Unlock()

	// Create a channel stream reader
	sr := channelreader.NewChannelStreamReader(d.ch[m])
	d.m[m] = sr

	return sr, nil
}

// Stop implements dialer.Dialer.
func (d *mockDialer) Close(m *types.Member) error {
	d.mx.Lock()
	defer d.mx.Unlock()

	delete(d.m, m)
	close(d.ch[m])
	delete(d.ch, m)

	return nil
}

// StopAll implements dialer.Dialer.
func (d *mockDialer) CloseAll() {
	d.mx.Lock()
	defer d.mx.Unlock()

	for _, v := range d.ch {
		close(v)
	}
	d.m = make(map[*types.Member]io.StreamReader)
	d.ch = make(map[*types.Member]chan *types.Message)
}
