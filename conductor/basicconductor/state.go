package basicconductor

type state byte

const (
	stateInit state = iota	// the goroutine is being initiated
	stateConn				// the goroutine is trying to connect with the member
	stateWork				// the goroutine has entered the loop in collector.Work()
)
