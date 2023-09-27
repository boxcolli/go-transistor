package pebblewal

import (
	"time"

	"github.com/boxcolli/pepperlink/wal"
	"github.com/cockroachdb/pebble"
)

type pebbleWAL struct {
	db *pebble.DB
}

// Flush implements wal.WAL.
func (*pebbleWAL) Flush(timeCursor time.Time) error {
	panic("unimplemented")
}

// Push implements wal.WAL.
func (*pebbleWAL) Push(topic string, topicId []byte, msg []byte, timestamp time.Time) error {
	panic("unimplemented")
}

func NewPebbleWAL(db *pebble.DB) wal.WAL {
	return &pebbleWAL{
		db: db,
	}
}
