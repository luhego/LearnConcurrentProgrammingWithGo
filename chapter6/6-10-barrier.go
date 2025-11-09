package main

import (
	"fmt"
	"sync"
	"time"
)

// Barrier is a reusable synchronization primitive that blocks goroutines
// until a predefined number (size) have all called Wait. Once the last
// goroutine arrives, all are released simultaneously, and the barrier
// resets for potential reuse.
type Barrier struct {
	size      int        // Total number of goroutines required to trip the barrier
	waitCount int        // Number of goroutines currently waiting
	cond      *sync.Cond // Condition variable used to coordinate waiting and release
}

// NewBarrier creates and returns a new Barrier that will release
// goroutines once 'size' have called Wait.
func NewBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})
	return &Barrier{size, 0, condVar}
}

// Wait blocks the calling goroutine until the total number of waiting
// goroutines equals the barrier size. The last goroutine to arrive
// releases all waiting goroutines and resets the counter, allowing
// the barrier to be reused.
func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount++
	if b.waitCount < b.size {
		b.cond.Wait()
	} else {
		b.waitCount = 0
		b.cond.Broadcast()
	}
	b.cond.L.Unlock()
}

func workAndWait(name string, timeToWork int, barrier *Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go workAndWait("Red", 4, barrier)
	go workAndWait("Blue", 10, barrier)
	time.Sleep(100 * time.Second)
}
