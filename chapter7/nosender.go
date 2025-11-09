package main

import (
	"fmt"
	"time"
)

func sender(message chan string) {
	time.Sleep(5 * time.Second)
	fmt.Println("Sender slept for 5 seconds")
}

func main() {
	msgChannel := make(chan string)
	go sender(msgChannel)
	fmt.Println("Reading message from channel...")
	msg := <-msgChannel
	fmt.Println("Received:", msg)
}
