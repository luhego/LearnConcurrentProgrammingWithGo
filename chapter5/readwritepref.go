package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Reader's counter: How many reader goroutines are actively accessing the shared resources
 * Writer's waiting counter: How many writer goroutines are suspended waiting to access the shared resources
 * Writer active indicator: Flag that tells us if the resource is currently being updated by a writer goroutine.
 * Condition variable with mutex: Allows us to set various conditions on the preceding properties, suspending execution when the conditions are met.
 */
type ReadWriteMutex struct {
	readersCounter        int
	writersWaitingCounter int
	writerActive          bool
	cond                  *sync.Cond
}

func NewReadWriteMutex() *ReadWriteMutex {
	return &ReadWriteMutex{cond: sync.NewCond(&sync.Mutex{})}
}

/*
 * Attempt to acquire the read lock.
 * We prevent new readers, if there are writers waiting
 */
func (rw *ReadWriteMutex) ReadLock() {
	rw.cond.L.Lock()
	for rw.writersWaitingCounter > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.readersCounter++
	rw.cond.L.Unlock()

}

/*
 * Attempt to acquire the write lock.
 * Writers are blocked when there are readers or writers using the resource.
 */
func (rw *ReadWriteMutex) WriteLock() {
	rw.cond.L.Lock()
	rw.writersWaitingCounter++
	for rw.readersCounter > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.writerActive = true
	rw.writersWaitingCounter--
	rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.cond.L.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.cond.Broadcast()
	}

	rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.cond.L.Lock()
	rw.writerActive = false
	rw.cond.Broadcast()
	rw.cond.L.Unlock()
}

func main() {
	rwMutex := NewReadWriteMutex()
	for i := 0; i < 2; i++ {
		go func() {
			for {
				rwMutex.ReadLock()
				time.Sleep(1 * time.Second)
				fmt.Println("Read done")
				rwMutex.ReadUnlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	rwMutex.WriteLock()
	fmt.Println("Write finished")
}
