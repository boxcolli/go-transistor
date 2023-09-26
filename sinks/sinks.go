package sinks

import "time"

type Sink interface {
	Write(topic string, topicId []byte, msg interface{}, timestamp time.Time) error
	Delete(topic string, topicId []byte) error
}


type SinkOption struct {
	createTopicTable bool
}

type MessageConverter interface {
	MessageToSchema(topic string, msg interface{}) interface{}
 SchemaToMessage(topic string, schema interface{}) interface{}
}

/*
Topic validator:
	데이터를 저장할 때 토픽에 해당하는 테이블이 없으면 치명적이다.
	INSERT 하기 전에 테이블이 존재하는지 로컬 캐시에서 확인한다. 만약 없다면 데이터베이스에 직접 테이블을 생성하고 캐시를 업데이트한다. 캐시를 업데이트할 때는 mutex lock을 걸어야 한다.
	이 기능을 자동으로 넣을지 안넣을지 결정 가능
*/
type TopicValidator interface {
	FetchTopicTable(topic string) (bool, error)
	CreateTopicTable(topic string) error
}