package collector

import "github.com/boxcolli/go-transistor/types"

// CollectorManager manages Collectors.
// It is recommended to implement CollectorManager goroutine safe.
type CollectorManager interface {
	Load(mid types.MemberId, s Stream) <-chan error
	Unload(mid types.MemberId)
}
