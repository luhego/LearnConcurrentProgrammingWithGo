package main

import (
	"fmt"
	"time"
)

type StopCondition func(n int) bool

func generator(out chan<- int, quit <-chan struct{}) {
	defer close(out)
	next := 1
	for {
		select {
		case <-quit:
			fmt.Println("genrator: received quit signal")
			return
		case out <- next:
			next++
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func worker(id int, numbers <-chan int, quit chan<- struct{}, stop StopCondition) {
	for n := range numbers {
		fmt.Printf("worker %d received %d\n", id, n)

		// Worker-specific stopping logic
		if stop(n) {
			fmt.Printf("worker %d: stopping condition met, quitting...\n", id)
			quit <- struct{}{}
			return
		}
	}
}

func main() {
	numbers := make(chan int)
	quit := make(chan struct{})

	// Start generator
	go generator(numbers, quit)

	// Two workers, either may quit
	go worker(1, numbers, quit, func(n int) bool {
		return n >= 10
	})
	go worker(2, numbers, quit, func(n int) bool {
		return n >= 1000
	})

	// Block until a quit signal arrives
	<-quit
	fmt.Println("main: received quit, shutting down")

	// Let goroutines flush prints
	time.Sleep(200 * time.Millisecond)
}
