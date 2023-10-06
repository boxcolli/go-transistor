package collectors

// Collector is a set of incoming streams from a cluster.
type Collector interface {
	// Load new stream and create a new goroutine that continuously collects data from stream
	Load(stream StreamAdapter)
}
