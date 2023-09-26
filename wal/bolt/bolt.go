package bolt

import (
	"github.com/boxcolli/pepperlink/wal"
	"time"
)

type boltWAL struct {
}

// Flush implements wal.WAL.
func (*boltWAL) Flush(timeCursor time.Time) error {
	panic("unimplemented")
}

// Push implements wal.WAL.
func (*boltWAL) Push(topic string, topicId []byte, msg []byte, timestamp time.Time) error {
	panic("unimplemented")
}

func NewBoltWAL() wal.WAL {
	return &boltWAL{}
}
