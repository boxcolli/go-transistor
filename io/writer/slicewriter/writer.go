package slicewriter

import (
	"sync"

	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type SliceWriter struct {
	M []*types.Message
	MX sync.Mutex
}

// Write implements io.StreamWriter.
func (w *SliceWriter) Write(m *types.Message) error {
	w.MX.Lock()
	w.M = append(w.M, m)
	w.MX.Unlock()
	return nil
}

func NewSliceWriter() (io.StreamWriter, *SliceWriter) {
	w := &SliceWriter{
		M: []*types.Message{},
	}
	return w, w
}
