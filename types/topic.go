package types

var (
	DefaultTopic = Topic{}
	DefaultTopics = []Topic{ DefaultTopic }
)

type Topic []string

// Empty topic means a wildcard for all kinds of topic
func (t Topic) Empty() bool {
	return len(t) == 0
}
