package types

import (
	"fmt"
	"time"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Message struct {
	Topic	Topic
	Mode	Mode
	Method 	Method
	Data 	[]byte
	TP 		time.Time
}

func (m *Message) Marshal() *pb.Message {
	return &pb.Message{
		Topic: 		&pb.Topic{ Tokens: m.Topic },
		Mode:		m.Mode.ToBuf(),
		Method:		m.Method.ToBuf(),
		Data:		&anypb.Any{ Value: m.Data },
		Timestamp:	timestamppb.New(m.TP),
	}
}

func (m *Message) Unmarshal(msg *pb.Message) error {

	m.Topic = msg.Topic.GetTokens()
	m.Mode.FromBuf(msg.GetMode())
	m.Method.FromBuf(msg.Method)
	m.Data = msg.Data.GetValue()
	m.TP = msg.Timestamp.AsTime()

	return nil
}

func (m Message) String() string {
	return fmt.Sprintf(
		"Topic%v Mode[%s] Method[%s] Data[%v] TP[%v]\n",
		m.Topic, m.Mode.String(), m.Method.String(), m.Data, m.TP,
	)
}
