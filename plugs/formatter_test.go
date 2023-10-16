package plugs

import (
	"testing"

	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

func TestBasicFormatter(t *testing.T) {
	prefix := "test"
	delimiter := "#"
	f := NewBasicFormatter(prefix, delimiter)

	member := types.Member{
		Cname: "c0",
		Name: "n0",
		Pro: types.ProtocolGrpc,
		Host: "localhost",
		Port: "443",
	}

	keyspace := prefix + delimiter + member.Cname + delimiter
	key := prefix + delimiter + member.Cname + delimiter + member.Name
	value := string(member.Pro) + delimiter + member.Host + delimiter + member.Port

	assert.Equal(t, keyspace, f.PrintKeyspace(member.Cname))
	assert.Equal(t, key, f.PrintKey(&member))
	assert.Equal(t, value, f.PrintValue(&member))

	{
		m := types.Member{}
		f.ScanKey(key, &m)
		assert.Equal(t, member.Cname, m.Cname)
		assert.Equal(t, member.Name, m.Name)
	}

	{
		m := types.Member{}
		f.ScanValue(value, &m)
		assert.Equal(t, member.Pro, m.Pro)
		assert.Equal(t, member.Host, m.Host)
		assert.Equal(t, member.Port, m.Port)
	}
}
