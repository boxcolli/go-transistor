package channelwriter

import (
	"log"
	"sync"
	"testing"

	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

var ms = []*types.Message{
	{ Topic: types.Topic{"A0"} },
	{ Topic: types.Topic{"A1"} },
}

func TestChannelWriter(t *testing.T) {
	c := make(chan *types.Message, 2)
	w := NewChannelWriter(c)

	for _, m := range ms { w.Write(m) }
	assert.Equal(t, len(ms), len(c))

	for i := 0; i < len(ms); i++ {
		log.Println("received:", *<-c)
	}
}

func TestConcurrency(t *testing.T) {
	c := make(chan *types.Message, 2)
	w := NewChannelWriter(c)

	start := make(chan bool)
	wg := sync.WaitGroup{}
	for _, m := range ms {
		wg.Add(1)
		go func(m *types.Message) {
			<- start
			w.Write(m)
			wg.Done()
		} (m)
	}

	close(start)
	wg.Wait()
	assert.Equal(t, len(ms), len(c))

	for i := 0; i < len(ms); i++ {
		log.Println("received:", *<-c)
	}
}
