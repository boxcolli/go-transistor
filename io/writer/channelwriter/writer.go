package channelwriter

import (
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type channelStreamWriter struct {
	ch	chan<- *types.Message
}

func (w *channelStreamWriter) Write(m *types.Message) error {
	w.ch <- m
	return nil
}

func NewChannelWriter(ch chan<- *types.Message) io.StreamWriter {
	return &channelStreamWriter{
		ch: ch,
	}
}
