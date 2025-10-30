package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
User calls:

	rw.ReadLock()
	    ↓
	    readersLock (short, for bookkeeping)
	    readersCounter++
	    if first → globalLock.Lock()
	    readersLock.Unlock()

	rw.ReadUnlock()
	    ↓
	    readersLock (short, for bookkeeping)
	    readersCounter--
	    if last → globalLock.Unlock()
	    readersLock.Unlock()

	rw.WriteLock()
	    ↓
	    globalLock.Lock()  ← mutual exclusion with readers

	rw.WriteUnlock()
	    ↓
	    globalLock.Unlock()
*/
type ReadWriteMutex struct {
	readersCounter int

	readersLock sync.Mutex

	globalLock sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock() // Synchronizes access so thay only one goroutine is allowed at any time
	rw.readersCounter++   // Reader goroutine increments readersCounter by 1
	if rw.readersCounter == 1 {
		rw.globalLock.Lock() // If a reader goroutine is the first one, it attempts to lock globaLock
	}
	rw.readersLock.Unlock() // Synchronizes access so that only one gorotuine is allowed at any time

}
func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock() // Any writer access requires a lock on globalLock
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()       // Synchronizes access so that only one goroutine is allowed at any time
	rw.readersCounter--         // The reader goroutine decrements readersCounter by 1
	if rw.readersCounter == 0 { // If the reader goroutine is the last one out, it unlocks the global lock
		rw.globalLock.Unlock() // Synchronizes access so that only one goroutine is allowed at any time
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
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
