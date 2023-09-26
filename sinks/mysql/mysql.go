package mysql

import (
	"github.com/boxcolli/pepperlink/sinks"
	"time"
)

type mysqlSink struct {
	opt sinks.SinkOption
}

// Write implements sinks.Sink.
func (s *mysqlSink) Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error {
	panic("unimplemented")
}

// Delete implements sinks.Sink.
func (s *mysqlSink) Delete(topic string, topicId []byte) error {
	panic("unimplemented")
}

func NewMySQLSink(opt sinks.SinkOption) sinks.Sink {
	return &mysqlSink{}
}