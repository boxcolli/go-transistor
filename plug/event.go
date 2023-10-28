package plug

import (
	"github.com/boxcolli/go-transistor/types"
)

type Event struct {
	Op		types.Operation
	Data	*types.Member
}
