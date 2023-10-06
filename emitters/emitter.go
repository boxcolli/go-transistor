package emitters

import (
	"time"

	pb "github.com/boxcolli/pepperlink/idl/gen/transistor/v1"
)

type Emitter interface {
	Emit(topic string, topicId []byte, method pb.Method, data interface{}, tp time.Time)
}
