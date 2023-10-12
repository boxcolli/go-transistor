package channel

import (
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type channelReader struct {
	ch chan *types.Message
}

// Read implements io.StreamReader.
func (r *channelReader) Read() (*types.Message, error) {
	m, ok := <-(r.ch)
	if !ok {
		return nil, io.ErrClosed
	}
	return m, nil
}

func NewChannelReader(ch chan *types.Message) io.StreamReader {
	return &channelReader{
		ch: ch,
	}
}
