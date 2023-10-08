package collector

import "github.com/boxcolli/go-transistor/io/reader"


type Collector interface {
	Work(r *reader.StreamReader) error
	Stop(r *reader.StreamReader)
	StopAll()
}
