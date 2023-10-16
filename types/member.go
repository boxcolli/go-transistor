package types

type Protocol byte

func (p Protocol) String() string {
	return string(p) // not Ascii value, same bit value as byte.
}

const (
	ProtocolUnspecified Protocol = iota
	ProtocolGrpc
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
