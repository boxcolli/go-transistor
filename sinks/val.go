package sinks

import "database/sql"

/*
Topic validator:
	데이터를 저장할 때 토픽에 해당하는 테이블이 없으면 치명적이다.
	INSERT 하기 전에 테이블이 존재하는지 로컬 캐시에서 확인한다. 만약 없다면 데이터베이스에 직접 테이블을 생성하고 캐시를 업데이트한다. 캐시를 업데이트할 때는 mutex lock을 걸어야 한다.
	이 기능을 자동으로 넣을지 안넣을지 결정 가능
*/
type TopicTableValidator interface {
	FetchTopicTable(topic string) (bool, error)
	CreateTopicTable(topic string) error
}

type SQLTopicTableValidator struct {
	db *sql.DB
}

// FetchTopicTable implements sinks.TopicTableValidator.
func (*SQLTopicTableValidator) FetchTopicTable(topic string) (bool, error) {
	panic("unimplemented")
}

// CreateTopicTable implements sinks.TopicTableValidator.
func (*SQLTopicTableValidator) CreateTopicTable(topic string) error {
	panic("unimplemented")
}

func NewSQLTopicTableValidator(db *sql.DB) *SQLTopicTableValidator {
	return &SQLTopicTableValidator{
		db: db,
	}
}
