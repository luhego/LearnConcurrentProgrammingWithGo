package main

import (
	"fmt"
	"time"
)

// A race condition happens when the value at index 100 is never updated and remains with value zero(or maybe it is updated but with value zero)
func addNextNumber(nextNum *[101]int) {
	i := 0
	// Skip indices that have been updated
	for nextNum[i] != 0 {
		i++
	}
	// Update current index
	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	nextNum := [101]int{1}
	// Creating 100 goroutines
	for i := 0; i < 100; i++ {
		go addNextNumber(&nextNum)
	}

	// The initial value of nextNum[100] is zero, we should exit this loop as soon as this value is updated
	for nextNum[100] == 0 {
		println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println(nextNum)
}
