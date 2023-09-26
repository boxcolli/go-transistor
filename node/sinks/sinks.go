package sinks

type Sink interface {
	func Write(topic string, msg interface{}, timestamp Time.time) error
	func Delete(topic string, id []byte) error
}
