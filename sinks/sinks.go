package sinks

import "time"

// 
type Sink interface {
	Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error
	Delete(topic string, topicId []byte) error
}

//
type SinkOption struct {
	validateTopicTable bool
}

