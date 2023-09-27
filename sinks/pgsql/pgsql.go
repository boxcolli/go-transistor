package pgsql

import (
	"time"
	"database/sql"

	"github.com/boxcolli/pepperlink/sinks"
)

type pgsqlSink struct {
	db		*sql.DB
	conv	sinks.MessageConverter
	val		sinks.TopicTableValidator
	opt		sinks.SinkOption
}
   
// Write implements sinks.Sink.
func (s *pgsqlSink) Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error {
	panic("unimplemented")
}

// Delete implements sinks.Sink.
func (s *pgsqlSink) Delete(topic string, topicId []byte) error {
	panic("unimplemented")
}

func NewPgSQLSink(db *sql.DB, conv sinks.MessageConverter, val sinks.TopicTableValidator, opt sinks.SinkOption) sinks.Sink {
	return &pgsqlSink{
		db: db,
		conv: conv,
		val: val,
		opt: opt,
	}
}
