package tools

import "time"

type Timer interface {
	Set()				// set another new timer session
	End() <-chan bool	// get asynchronous notification
	Stop()				// release the goroutine
}

type timer struct {
	d   	time.Duration
	set 	chan bool
	end 	chan bool
	stop	chan bool
}

// Warning: the goroutine does not automatically stop.
// Timer.Stop() should be called manually.
func NewTimer(d time.Duration) Timer {
	t := &timer{
		d:   d,
		set: make(chan bool, 1),
		end: make(chan bool, 1),
		stop: make(chan bool, 1),
	}
	go t.run()

	return t
}

func (t *timer) run() {
	for {
		select {
		case <- t.stop:	// prioritize stop signal
			return

		case <- t.set:
			time.Sleep(t.d)
			t.end <- true
		}
	}
}

func (t *timer) Set() {
	t.set <- true
}

func (t *timer) End() <-chan bool {
	return t.end
}

func (t *timer) Stop() {
	t.stop <- true
}
