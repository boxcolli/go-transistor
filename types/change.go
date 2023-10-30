package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

var (
	DefaultChange = Change{
		Op: OperationAdd,
		Topic: DefaultTopic,
	}
)

type Change struct {
	Op		Operation
	Topic	Topic
}

func (m *Change) Marshal() *pb.Change {
	return &pb.Change{
		
	}
}

func (m *Change) Unmarshal(msg *pb.Change) {

}
