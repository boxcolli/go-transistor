package collector

import "github.com/boxcolli/go-transistor/base"

/*
	Collector is a wrapper of Stream and collects messages from it.
	Stream is not goroutine safe, so there should be only one goroutine
	running with Collector instance.
*/
type Collector interface {
	// Start() starts a goroutine that collects messages from stream
	// and passes them into Base.
	Start(base *base.Base) <-chan error

	// Stop() stops the goroutine of Start()
	Stop()
}
