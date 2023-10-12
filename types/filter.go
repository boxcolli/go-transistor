package types

type Filter struct {
	// A set of topic
	// If nil, it's wildcard - receive all topics.
	T []Topic
}
