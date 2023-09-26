package bolt

import (
	"github.com/boxcolli/pepperlink/sinks"
	"time"
)

type boltSink struct {
	opt sinks.SinkOption
	
}

// Write implements sinks.Sink.
func (s *boltSink) Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error {
	panic("unimplemented")
}

// Delete implements sinks.Sink.
func (s *boltSink) Delete(topic string, topicId []byte) error {
	panic("unimplemented")
}

func NewBoltSink(opt sinks.SinkOption) sinks.Sink {
	return &boltSink{
		opt: opt,
	}
}
