package plugs

const (
	KeyCid = "cid"
	EmptyCid = "emp"
)

type Plug interface {
	WatchPub()
	Destroy()
}