package boltsink

import (
	"github.com/boxcolli/pepperlink/sinks"
)

type boltMessageConverter struct {
	/*
		mtos:
			key: topic
			value: a function that converts protobuf message struct to string
	*/
	mtos map[string]func(interface{}) ([]byte, error)
}

func (m *boltMessageConverter) MessageToSchema(topic string, msg interface{}) ([]byte, error) {
	// receive protobuf message, convert it into string
	return m.mtos[topic](msg)
	// panic("unimplemented")
}

func NewBoltMessageConverter(mtos map[string]func(interface{}) ([]byte, error)) sinks.MessageConverter[interface{}, []byte] {
	return &boltMessageConverter{
		mtos: mtos,
	}
}