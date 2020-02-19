package util

import "time"

type Timer struct {
	start int64
	end   int64
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	t.start = time.Now().UnixNano()
}

func (t *Timer) Stop() {
	t.end = time.Now().UnixNano()
}

func (t *Timer) PrintElapsed(message string) {
	// elapsed := time.Now().Sub(start)
}
