// This program should continuosly consume from output channel, print the ouput and stop after 5 seconds

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers() chan int {
	output := make(chan int)
	go func() {
		for {
			output <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func main() {
	output := generateNumbers()

	stop := time.After(5 * time.Second)
	for output != nil {
		select {
		case tNow := <-stop:
			fmt.Println("Timed out. Waited until:", tNow.Format("15:04:05"))
			output = nil
		case msg := <-output:
			fmt.Println("Message received:", msg)
		}
	}
}
