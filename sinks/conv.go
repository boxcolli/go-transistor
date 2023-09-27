package sinks

type MessageConverter[M any, S any] interface {
	MessageToSchema(topic string, msg M) (S, error)
}

type SqlMessageConverter struct {
	/*
		mtos:
			key: topic
			value: a function that converts protobuf message struct to string
	*/
	mtos map[string]func(interface{}) (string, error)
}

func (m *SqlMessageConverter) MessageToSchema(topic string, msg interface{}) (string, error) {
	// receive protobuf message, convert it into string
	return m.mtos[topic](msg)
}

func NewSQLMessageConverter(mtos map[string]func(interface{}) (string, error)) *SqlMessageConverter {
	return &SqlMessageConverter{
		mtos: mtos,
	}
}