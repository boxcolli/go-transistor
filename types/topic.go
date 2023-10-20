package types

const (
	TopicWildcard = "*"
)

var (
	EmptyTopic = Topic{}
	DefaultTopic = Topic{ TopicWildcard }
	DefaultTopics = []Topic{ DefaultTopic }
)

type Topic []string

func (t Topic) Empty() bool {
	return len(t) == 0
}
