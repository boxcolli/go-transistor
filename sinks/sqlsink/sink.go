package sqlsink

import (
	"database/sql"
	"time"

	"github.com/boxcolli/go-transistor/sinks"
)

type Query struct {
	DB *sql.DB

	CS map[string]sql.Stmt
	US map[string]sql.Stmt
	DS map[string]sql.Stmt
	
	// data must be a struct pointer
	Create  map[string]func(topicId []byte, data interface{}, tp time.Time) error
	Update	map[string]func(topicId []byte, data interface{}, tp time.Time) error
	Delete	map[string]func(topicId []byte, tp time.Time) error
}

type sqlSink struct {
	query	Query
	opt    	sinks.SinkOption
}

func NewSQLSink(query Query, opt sinks.SinkOption) sinks.Sink {
	return &sqlSink{
		query: query,
		opt:  opt,
	}
}

// Create implements sinks.Sink.
func (s *sqlSink) Create(topic string, topicId []byte, data interface{}, tp time.Time) error {
	return s.query.Create[topic](topicId, data, tp)
}

// Update implements sinks.Sink.
func (s *sqlSink) Update(topic string, topicId []byte, data interface{}, tp time.Time) error {
	return s.query.Update[topic](topicId, data, tp)
}

// Delete implements sinks.Sink.
func (s *sqlSink) Delete(topic string, topicId []byte, tp time.Time) error {
	return s.query.Delete[topic](topicId, tp)
}
