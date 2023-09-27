package bolt

import (
	"testing"

	pb "github.com/boxcolli/pepperlink/idl/gen/hello/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	// "github.com/stretchr/testify/assert"
)

func TestBoltWAL(t *testing.T) {
	msg := pb.Hello{
		Name: "Alice",
	}
	t.Log(protojson.Marshal(&msg)) // serialization: protobuf => json string

	msgByte, _ := proto.Marshal(&msg) // serialization: protobuf => []byte
	t.Log(msgByte)
}