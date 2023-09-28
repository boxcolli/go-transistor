package sinks

import "time"

type Converter[M any, S any] interface {
	GetWriteQuery(topic string, topicId []byte, msg M, timestamp time.Time) (S, error)
	GetDeleteQuery(topic string, topicId []byte, msg M, timestamp time.Time) (S, error)
}

type MTOW map[string]func([]byte, interface{}, time.Time) (string, error)
type MTOD map[string]func([]byte, interface{}, time.Time) (string, error)

type SqlMessageConverter struct {
	/*
		mtos:
			key: topic
			value: a function that converts protobuf message struct to string
	*/
	mtow MTOW
	mtod MTOD
}

func (m *SqlMessageConverter) GetWriteQuery(topic string, topicId []byte, msg interface{}, timestamp time.Time) (string, error) {
	// receive protobuf message, convert it into string
	return m.mtow[topic](topicId, msg, timestamp)
}

func (m *SqlMessageConverter) GetDeleteQuery(topic string, topicId []byte, msg interface{}, timestamp time.Time) (string, error) {
	// receive protobuf message, convert it into string
	return m.mtod[topic](topicId, msg, timestamp)
}

func NewSQLMessageConverter(mtow MTOW, mtod MTOD) *SqlMessageConverter {
	return &SqlMessageConverter{
		mtow: mtow,
		mtod: mtod,
	}
}
