package types

import pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"

type Change struct {
	Op		Operation
	Topic	Topic
}

func NewChange(op Operation, t Topic) *Change {
	return &Change{ op, t }
}

func (c *Change) Marshal() *pb.Change {
	return &pb.Change{
		Op: c.Op.ToBuf(),
		Topic: &pb.Topic{ Tokens: c.Topic },
	}
}

func (c *Change) Unmarshal(cg *pb.Change) {
	c.Op.FromBuf(cg.Op)
	c.Topic = cg.Topic.GetTokens()
}
