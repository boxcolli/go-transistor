package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

type Change struct {
	Op		Operation
	Topic	Topic
}

func NewChange(op Operation, t Topic) *Change {
	return &Change{ op, t }
}

func (m *Change) Marshal() *pb.Change {
	return &pb.Change{
		
	}
}

func (m *Change) Unmarshal(msg *pb.Change) {

}
