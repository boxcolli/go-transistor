package plugs

import (
	"strings"

	"github.com/boxcolli/go-transistor/types"
)

type Formatter interface {
	PrintKeyspace(cname string) string
	PrintKey(m *types.Member) string
	PrintValue(m *types.Member) string

	ScanKey(key string, m *types.Member)
	ScanValue(value string, m *types.Member)
}

/*
	Key		_PREFIX_:_CNAME_:_NAME_
	Value	_PRO_:_HOST_:_PORT_
*/
type basicFormatter struct {
	prefix	string
	delim	string
	delimByte byte
}

func NewBasicFormatter(prefix string, delim string) Formatter {
	return &basicFormatter{
		prefix: prefix,
		delim: delim,
		delimByte: byte(delim[0]),
	}
}

// PrintKeyspace implements Formatter.
func (f *basicFormatter) PrintKeyspace(cname string) string {
	return f.prefix + f.delim + cname + f.delim
}

// PrintKey implements Formatter.
func (f *basicFormatter) PrintKey(m *types.Member) string {
	return f.prefix + f.delim + m.Cname + f.delim + m.Name
}

// PrintValue implements Formatter.
func (f *basicFormatter) PrintValue(m *types.Member) string {
	return m.Pro.String() + f.delim + m.Host + f.delim + m.Port
}

// ScanKey implements Formatter.
func (f *basicFormatter) ScanKey(key string, m *types.Member) {
	tokens := strings.Split(key, f.delim)
	m.Cname = tokens[1]
	m.Name = tokens[2]
}

// ScanValue implements Formatter.
func (f *basicFormatter) ScanValue(value string, m *types.Member) {
	tokens := strings.Split(value, f.delim)
	m.Pro = types.Protocol(tokens[0])
	m.Host = tokens[1]
	m.Port = tokens[2]
}
