package collector

import "github.com/boxcolli/go-transistor/io"

type Collector interface {
	Work(r io.StreamReader, call func(e error)) //error
	Stop(r io.StreamReader)
	StopAll()
}
