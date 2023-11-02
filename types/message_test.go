package types

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	var text = []byte("hello")
	var tp = []string{"1"}
	tn := time.Now().UTC()

	m := Message{

		Topic:  tp,
		Method: 0,
		Data:   text,
		TP:     tn,
	}
	newMessage := Message{}

	proto := m.Marshal()        // message -> protobuf
	newMessage.Unmarshal(proto) //protobuf -> message

	//assert.Equal(t, m, newMessage)
	assert.Equal(t, true, reflect.DeepEqual(m, newMessage))
}
