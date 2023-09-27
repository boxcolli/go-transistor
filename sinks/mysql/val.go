package mysql

import (
	"database/sql"

	"github.com/boxcolli/pepperlink/sinks"
)

type topicTableValidator struct {
	db *sql.DB
}

// FetchTopicTable implements sinks.TopicTableValidator.
func (*topicTableValidator) FetchTopicTable(topic string) (bool, error) {
	panic("unimplemented")
}

// CreateTopicTable implements sinks.TopicTableValidator.
func (*topicTableValidator) CreateTopicTable(topic string) error {
	panic("unimplemented")
}

func NewTopicTableValidator(db *sql.DB) sinks.TopicTableValidator {
	return &topicTableValidator{
		db: db,
	}
}
