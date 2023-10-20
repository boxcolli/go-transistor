package conductors

// Conductor conducts information exchange between Plug, Dialer and Collector.
type Conductor interface {
	// Start member discovery on the cluster
	Begin(cname string)	error

	// Stop member discovery and disconnect with members on the cluster
	End(cname string)

	// Stop all member discoveries
	EndAll()
}
