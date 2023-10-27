package types

type Protocol string

func (p Protocol) String() string {
	return string(p) // not Ascii value, same bit value as byte.
}

const (
	ProtocolUnspecified = "nil"
	ProtocolGrpc = "grpc"
)

type Member struct {
	Cname	string
	Name 	string
	Pro		Protocol
	Host	string
	Port 	string
}

func (m Member) Address() string {
	return m.Host + ":" + m.Port
}

func (m Member) EqualsId(x Member) bool {
	if m.Cname == x.Cname && m.Name == x.Name {
		return true
	}
	return false
}

func (m Member) Equals(x Member) bool {
	if m.Cname == x.Cname &&
		m.Name == x.Name &&
		m.Pro == x.Pro &&
		m.Host == x.Host &&
		m.Port == x.Port {
		return true
	}
	return false
}
