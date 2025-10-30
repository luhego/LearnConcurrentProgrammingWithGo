package main

import (
	"fmt"
	"sync"
	"time"
)

type ReadWriteMutex struct {
	readersCounter int

	readersLock sync.Mutex

	globalLock sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()

}
func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

func (rw *ReadWriteMutex) TryLock() bool {
	return rw.globalLock.TryLock()
}

func (rw *ReadWriteMutex) TryReadLock() bool {
	return rw.readersLock.TryLock()
}

func main() {
	var rw ReadWriteMutex

	reader1 := func() {
		fmt.Println("R1: trying ReadLock")
		rw.ReadLock()
		fmt.Println("R1: acquired ReadLock")
		time.Sleep(500 * time.Millisecond)
		rw.ReadUnlock()
		fmt.Println("R1: released ReadLock")
	}

	reader2 := func() {
		fmt.Println("R2: trying ReadLock")
		rw.ReadLock()
		fmt.Println("R2: acquired ReadLock")
		time.Sleep(500 * time.Millisecond)
		rw.ReadUnlock()
		fmt.Println("R2: released ReadLock")
	}

	writer1 := func() {
		fmt.Println("W1: trying WriteLock")
		rw.WriteLock()
		fmt.Println("W1: acquired WriteLock")
		time.Sleep(500 * time.Millisecond)
		rw.WriteUnlock()
		fmt.Println("W1: released WriteLock")
	}

	// Launch the goroutines with slight staggered timing
	go reader1()
	time.Sleep(100 * time.Millisecond)
	go reader2()
	time.Sleep(200 * time.Millisecond)
	go writer1()

	time.Sleep(2 * time.Second)
	fmt.Println("Main exiting")
}
