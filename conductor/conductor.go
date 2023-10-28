package conductor

import "github.com/boxcolli/go-transistor/types"

// Conductor conducts information exchange between Plug, Dialer and Collector.
type Conductor interface {
	// Start member discovery on the cluster
	Begin(cname string) error

	// Apply global topic change
	Apply(cg *types.Change)

	// Stop member discovery and disconnect with members on the cluster
	End(cname string)

	// Stop all member discoveries
	EndAll()
}
