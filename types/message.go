package types

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

type Message struct {
	Topic		Topic
	Method 		Method
	Data 		[]byte
	TP 			time.Time
}

func (m *Message) Marshal() *pb.Message {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(m.Data); err != nil {
		log.Fatalf("Failed to encode interface{}: %v", err)
	}
	rawBytes := buf.Bytes()

	return &pb.Message{
		Topic: &pb.Topic{Tokens: m.Topic},
		Method: m.Method.ToBuf(),
		Data: &anypb.Any{
			Value: rawBytes,
		},
	}
}

func (m *Message) Unmarshal(msg *pb.Message) {

}
