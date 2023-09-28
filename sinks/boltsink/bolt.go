package boltsink

import (
	"time"

	"github.com/boxcolli/pepperlink/sinks"
	"github.com/boltdb/bolt"
)

type boltSink struct {
	dbs		map[string]*bolt.DB
	conv	sinks.Converter[interface{}, []byte]
	opt		sinks.SinkOption
}

// Write implements sinks.Sink.
func (s *boltSink) Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error {
	panic("unimplemented")
}

// Delete implements sinks.Sink.
func (s *boltSink) Delete(topic string, topicId []byte) error {
	panic("unimplemented")
}

func NewBoltSink(dbs map[string]*bolt.DB, conv sinks.Converter[interface{}, []byte], opt sinks.SinkOption) sinks.Sink {
	return &boltSink{
		dbs: dbs,
  		conv: conv,
		opt: opt,
	}
}
