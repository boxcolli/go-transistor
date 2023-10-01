package plugs

const (
	KeyCid = "cid"
	EmptyCid = "emp"
)

type Plug interface {
	// Return plug type and an address that another node can plug with.
	GetDiscoveryAddr() ()
	WatchPub()
	Destroy()
}