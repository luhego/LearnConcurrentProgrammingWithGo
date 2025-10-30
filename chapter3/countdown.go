package main

import (
	"fmt"
	"time"
)

func main() {
	count := 5           // Allocates memory space for an integer variable
	go countdown(&count) // Starts goroutine and shares memory at the variable reference
	for count > 0 {
		time.Sleep(500 * time.Millisecond) // main goroutine reads the value of the shared variable every half second
		fmt.Println(count)
	}
}

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1 // The goroutine updates the value of the shared variable
	}
}
