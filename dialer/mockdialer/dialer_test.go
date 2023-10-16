package mockdialer

import (
	"sync"
	"testing"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

const (
	maxMember = 100
	maxMessage = 10
)

func TestMockDialer(t *testing.T) {
	m := map[*types.Member]io.StreamReader{}
	ms := []*types.Member{}
	{
		for i := 0; i < maxMember; i++ {
			ms = append(ms, new(types.Member))
		}
		ms[0].Name = "0"
	}
	ch := map[*types.Member]chan *types.Message{}
	{
		for _, v := range ms {
			ch[v] = make(chan *types.Message, maxMessage)
		}
	}
	d := NewMockDialer(m, ch)

	// Run Dial() concurrently
	t.Log("Calling Dial() concurrently..")
	readCount := map[*types.Member]int{}
	readmx := sync.Mutex{}
	start := make(chan bool)
	done := make(chan bool, maxMember)
	for _, v := range ms {
		go func(m *types.Member) {
			// Dial()
			sr, err := d.Dial(m)
			if err != nil {
				return
			}

			done <- true
			
			// Waif for start
			if m.Name == "0" {
				t.Log("0: waiting for start")
			}
			<- start
			
			// StreamReader.Read()
			count := 0
			for {
				_, err := sr.Read()
				if err != nil {
					break
				}
				if m.Name == "0" {
					t.Log("0: Read()")
				}
				count++
			}
			
			// Update count
			readmx.Lock()
			readCount[m] = count
			readmx.Unlock()

			// Done
			if m.Name == "0" {
				t.Log("0: done")
			}
			done <- true
		} (v)
	}

	// Wait for Dial() done
	for i := 0; i < maxMember; i++ {
		<- done
	}

	// Push messages
	t.Log("Pushing messages..")
	msgs := []*types.Message{}
	for i := 0; i < maxMessage; i++ {
		msgs = append(msgs, new(types.Message))
	}
	for _, msg := range msgs {
		for _, c := range ch {
			c <- msg
		}
	}

	// Start goroutine to read
	close(start)

	// Close channels
	t.Log("Closing channels..")
	d.CloseAll()

	// Await goroutines
	t.Log("Waiting for goroutines to stop..")
	for i := 0; i < maxMember; i++ {
		<- done
	}

	// Assert channel closed
	for _, v := range ch {
		_, ok := <- v
		assert.False(t, ok)
	}

	// Assert count
	for _, count := range readCount {
		assert.Equal(t, maxMessage, count)
	}
}
