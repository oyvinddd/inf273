package util

import (
	"fmt"
	"time"
)

// Timer struct for timing program execution
type Timer struct {
	start time.Time
	end   time.Time
}

// NewTimer convenience constructor
func NewTimer() *Timer {
	return &Timer{
		start: time.Now(),
	}
}

// Start the timer
func (t *Timer) Start() {
	t.start = time.Now()
}

// PrintElapsed time to the terminal
func (t *Timer) PrintElapsed() {
	fmt.Printf("\n(Program execution took %v)\n", time.Since(t.start))
}
