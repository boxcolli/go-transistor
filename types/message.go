package types

import (
	"time"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Message struct {
	Topic		Topic
	Method 		Method
	Data 		[]byte
	TP 			time.Time
}

func (m *Message) Marshal() *pb.Message {

	ts := timestamppb.New(m.TP)

	return &pb.Message{
		Topic:  &pb.Topic{Tokens: m.Topic},
		Method: m.Method.ToBuf(),
		Data: &anypb.Any{
			Value: m.Data,
		},
		Timestamp: ts,
	}
}

func (m *Message) Unmarshal(msg *pb.Message) error {

	m.TP = msg.Timestamp.AsTime()
	m.Topic = msg.Topic.GetTokens()
	m.Method.ToPb(msg.Method)
	m.Data = msg.Data.GetValue()

	return nil
}
