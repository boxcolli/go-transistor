package mysql

import (
	"database/sql"
	"time"

	"github.com/boxcolli/pepperlink/sinks"
)

type mysqlSink struct {
	db		*sql.DB
	conv	sinks.MessageConverter
	val		sinks.TopicTableValidator
	opt		sinks.SinkOption
}

// Write implements sinks.Sink.
func (s *mysqlSink) Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error {
	panic("unimplemented")
}

// Delete implements sinks.Sink.
func (s *mysqlSink) Delete(topic string, topicId []byte) error {
	panic("unimplemented")
}

func NewMySQLSink(db *sql.DB, conv sinks.MessageConverter, val sinks.TopicTableValidator, opt sinks.SinkOption) sinks.Sink {
	return &mysqlSink{
		db: db,
		conv: conv,
		val: val,
		opt: opt,
 	}
}
