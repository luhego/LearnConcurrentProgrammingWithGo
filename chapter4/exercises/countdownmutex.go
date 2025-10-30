package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutex := sync.Mutex{}
	count := 5                   // Allocates memory space for an integer variable
	go countdown(&count, &mutex) // Starts goroutine and shares memory at the variable reference
	remaining := count
	for remaining > 0 {
		time.Sleep(500 * time.Millisecond) // main goroutine reads the value of the shared variable every half second
		mutex.Lock()
		fmt.Println(count)
		remaining = count
		mutex.Unlock()
	}
}

func countdown(seconds *int, mutex *sync.Mutex) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		*seconds -= 1 // The goroutine updates the value of the shared variable
		mutex.Unlock()
	}
}
