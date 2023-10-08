package types

import (
	"time"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
)

type Message struct {
	Topic		string
	Subtopic	string
	Method 		Method
	Data 		interface{}
	TP 			time.Time
}

func (m *Message) Marshal() *pb.Message {
	return &pb.Message{

	}
}

func (m *Message) Unmarshal(msg *pb.Message) {

}
