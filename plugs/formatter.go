package plugs

import (
	"github.com/boxcolli/go-transistor/types"
)

type Formatter interface {
	PrintKeyspace(cname string) string
	PrintKey(m *types.Member) string
	PrintValue(m *types.Member) string

	ScanKey(key []byte, m *types.Member)
	ScanValue(value []byte, m *types.Member)
}

type basicFormatter struct {
	prefix string
}

func NewBasicFormatter(prefix string) Formatter {
	return &basicFormatter{prefix: prefix}
}

/*
	Key		_PREFIX_:_CNAME_:_NAME_
	Value	_PRO_:_HOST_:_PORT_
*/

const (
	delimiter = "#"
	delimiterByte = byte('#')
)

// PrintKeyspace implements Formatter.
func (f *basicFormatter) PrintKeyspace(cname string) string {
	return f.prefix + delimiter + cname + delimiter
}

// PrintKey implements Formatter.
func (f *basicFormatter) PrintKey(m *types.Member) string {
	return f.prefix + delimiter + m.Cname + delimiter + m.Name
}

// PrintValue implements Formatter.
func (*basicFormatter) PrintValue(m *types.Member) string {
	return m.Pro.String() + delimiter + m.Host + delimiter + m.Port
}

// ScanKey implements Formatter.
func (*basicFormatter) ScanKey(key []byte, m *types.Member) {
	lo, hi := 0, 0
	hi = next(key, hi) // discard prefix
	
	lo = hi + 1
	hi = next(key, hi + 1)
	m.Cname = string(key[lo:hi])

	lo = hi + 1
	hi = next(key, hi + 1)
	m.Name = string(key[lo:hi])
}

// ScanValue implements Formatter.
func (*basicFormatter) ScanValue(value []byte, m *types.Member) {
	lo, hi := 0, 0
	hi = next(value, hi)
	m.Pro = types.Protocol(value[lo])

	lo = hi + 1
	hi = next(value, hi + 1)
	m.Host = string(value[lo:hi])

	lo = hi + 1
	hi = next(value, hi + 1)
	m.Port = string(value[lo:hi])
}

func next(b []byte, i int) int {
	for ; i < len(b); i++ {
		if b[i] == delimiterByte {
			return i
		}
	}
	return i
}
