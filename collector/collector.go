package collector

import "github.com/boxcolli/go-transistor/io"

type Collector interface {
	Work(r io.StreamReader) error
	Stop(r io.StreamReader)
	StopAll()
}
