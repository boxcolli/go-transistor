package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMember(t *testing.T) {
	x := Member{
		Cname: "c0",
		Name:  "n0",
		Pro:   ProtocolUnspecified,
	}

	y := Member{
		Cname: "c0",
		Name:  "n0",
		Pro:   ProtocolGrpc,
	}

	assert.False(t, x.Equals(y))
	assert.True(t, x.EqualsId(y))
}
