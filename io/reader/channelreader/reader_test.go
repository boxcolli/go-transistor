package channelreader

import (
	"testing"

	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

var ms = []*types.Message{
	{ Topic: types.Topic{"0"} },
	{ Topic: types.Topic{"1"} },
	{ Topic: types.Topic{"2"} },
}

func TestChannelStreamReader(t *testing.T) {
	ch := make(chan *types.Message, 10)
	sr := NewChannelStreamReader(ch)

	go func(ms []*types.Message) {
		for _, v := range ms {
			ch <- v
		}
		close(ch)
	} (ms)

	count := 0
	for {
		m, err := sr.Read()
		if err != nil {
			t.Log("channel end:", err)
			break
		}

		count++
		t.Log("Read():", m)
	}

	assert.Equal(t, len(ms), count)
}
