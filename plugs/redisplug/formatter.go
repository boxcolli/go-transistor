package redisplug

import (
	"fmt"
	"strings"

	"github.com/boxcolli/go-transistor/plugs"
	"github.com/boxcolli/go-transistor/types"
)

/*
	Keyspace	__keyspace@<DB>__:_PREFIX_:_CNAME_:*
	Key			_PREFIX_:_CNAME_:_NAME_
	Value		_PRO_:_HOST_:_PORT_
*/

type RedisFormatter interface {
	plugs.Formatter
	PrintPSubscribeKeyspace(cname string) string
}

type redisFormatter struct {
	channel	string
	prefix	string
	delim	string
}

func NewBasicRedisFormatter(dbnum, prefix, delim string) RedisFormatter {
	return &redisFormatter{
		channel: fmt.Sprintf("__keyspace@%s__:", dbnum),
		prefix: prefix,
		delim: delim,
	}
}

func (f *redisFormatter) PrintPSubscribeKeyspace(cname string) string {
	return f.channel + f.prefix + f.delim + cname + f.delim + "*"
}

// PrintKeyspace implements plugs.Formatter.
func (f *redisFormatter) PrintKeyspace(cname string) string {
	return f.prefix + f.delim + cname + f.delim + "*"
}

// PrintKey implements plugs.Formatter.
func (f *redisFormatter) PrintKey(m *types.Member) string {
	return f.prefix + f.delim + m.Cname + f.delim + m.Name
}

// PrintValue implements plugs.Formatter.
func (f *redisFormatter) PrintValue(m *types.Member) string {
	return m.Pro.String() + f.delim + m.Host + f.delim + m.Port
}

// ScanKey implements plugs.Formatter.
func (f *redisFormatter) ScanKey(key string, m *types.Member) {
	tokens := strings.Split(key, f.delim)
	m.Cname = tokens[1]
	m.Name = tokens[2]
}

// ScanValue implements Formatter.
func (f *redisFormatter) ScanValue(value string, m *types.Member) {
	tokens := strings.Split(value, f.delim)
	m.Pro = types.Protocol(tokens[0])
	m.Host = tokens[1]
	m.Port = tokens[2]
}
