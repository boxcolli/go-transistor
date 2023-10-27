package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

var (
	DefaultChange = Change{
		Op: OperationAdd,
		Topics: DefaultTopics,
	}
)

type Change struct {
	Op		Operation
	Topics	[]Topic
}

func (m *Change) Marshal() *pb.Change {
	return &pb.Change{
		
	}
}

func (m *Change) Unmarshal(msg *pb.Change) {

}
