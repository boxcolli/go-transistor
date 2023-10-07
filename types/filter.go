package types

type Filter struct {
	// A set of topic
	// If nil, it's wildcard - receive all topics.
	T map[string]bool

	// A set of topic id
	// If nil, it's wildcard - receive all topic ids
	Tid map[string]bool
}
