package bolt

import (
	"time"

	"github.com/boxcolli/pepperlink/wal"
	"github.com/boltdb/bolt"
)

type boltWAL struct {
	db *bolt.DB
}

// Flush implements wal.WAL.
func (*boltWAL) Flush(timeCursor time.Time) error {
	panic("unimplemented")
}

// Push implements wal.WAL.
func (*boltWAL) Push(topic string, topicId []byte, msg []byte, timestamp time.Time) error {
	panic("unimplemented")
}

func NewBoltWAL(db *bolt.DB) wal.WAL {
	return &boltWAL{
		db: db,
	}
}
