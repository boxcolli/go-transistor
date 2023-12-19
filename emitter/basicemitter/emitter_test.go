package basicemitter

import (
	"testing"
	"time"

	"github.com/boxcolli/go-transistor/io/writer/channelwriter"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

const (
	qsiz = 10
)

var ms = []*types.Message{
	{Topic: types.Topic{"A0"}},
	{Topic: types.Topic{"A1"}},
}

func TestBasicEmitter(t *testing.T) {
	emitter := NewBasicEmitter(qsiz)
	messageChannel := make(chan *types.Message, qsiz)
	streamWriter := channelwriter.NewChannelWriter(messageChannel)

	// Work
	go func() {
		err := emitter.Work(streamWriter)
		assert.NoError(t, err)
	} ()
	time.Sleep(500 * time.Millisecond)

	bus, ok := emitter.Bus(streamWriter)
	assert.True(t, ok)

	// Read message
	go func() {
		for {
			m, ok := <- messageChannel
			assert.True(t, ok)
			t.Logf("received: %v\n", m)
		}
	} ()

	bus.Push(ms[0])
	time.Sleep(500 * time.Millisecond)

	bus.Push(ms[1])
	time.Sleep(1000 * time.Millisecond)

}
