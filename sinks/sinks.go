package sinks

import "time"

// 
type Sink interface {
	Create(topic string, topicId []byte, data interface{}, tp time.Time) error
	Update(topic string, topicId []byte, data interface{}, tp time.Time) error
	Delete(topic string, topicId []byte, tp time.Time) error
}

type SinkOption struct {
	// validateTopicTable	bool
}
