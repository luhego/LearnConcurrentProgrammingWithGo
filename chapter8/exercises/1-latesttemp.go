package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulates reading and sending the temperature every 200ms
func generateTemp() chan int {
	output := make(chan int)
	go func() {
		temp := 50 // fahrenheit
		for {
			output <- temp
			temp += rand.Intn(3) - 1
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

// Outpus a message found in the input channel every 2 seconds
func outputTemp(input chan int) {
	go func() {
		for {
			fmt.Println("Current temp:", <-input)
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	inputChannel := make(chan int)
	outputTemp(inputChannel)
	tempChannel := generateTemp()

	temp := <-tempChannel
	for {
		select {
		case temp = <-tempChannel:
		case inputChannel <- temp:
		}
	}
}
