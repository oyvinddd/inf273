package util

import (
	"fmt"
	"time"
)

type Timer struct {
	start time.Time
	end   time.Time
}

func NewTimer() Timer {
	return Timer{}
}

func (t *Timer) Start() {
	t.start = time.Now() //.UnixNano()
}

func (t *Timer) Stop() {
	t.end = time.Now() //.UnixNano()
}

func (t *Timer) PrintElapsed() {
	fmt.Printf("\nProgram execution took %v", time.Since(t.start))
}
