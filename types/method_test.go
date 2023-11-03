package types

import (
	"testing"

	pb "github.com/boxcolli/go-transistor/api/gen/transistor/v1"
	"github.com/stretchr/testify/assert"
)

func TestMethod(t *testing.T) {
	m := MethodCreate
	toBuf := m.ToBuf()

	buf := pb.Method_METHOD_CREATE
	var fromBuf Method
	fromBuf.FromBuf(buf)

	assert.Equal(t, pb.Method_METHOD_CREATE, toBuf, "toBuf")
	assert.Equal(t, MethodCreate, fromBuf, "fromBuf")
}