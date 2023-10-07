package dialer

import "github.com/boxcolli/go-transistor/types"

type Dialer interface {
	Dial(m types.Member)
	Add(m types.Member)
	Delete(m types.Member)
}
