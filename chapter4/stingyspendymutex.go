package main

import (
	"fmt"
	"sync"
	"time"
)

/*
* Adds 10 dollars 1 million times
 */
func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}
	fmt.Println("Stingy Done")
}

/*
* Substracts 10 dollars 1 million times
 */
func spendy(money *int, mutext *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		mutext.Lock()
		*money -= 10
		mutext.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100

	mutex := sync.Mutex{}

	go stingy(&money, &mutex)
	go spendy(&money, &mutex)

	time.Sleep(2 * time.Second)

	mutex.Lock()
	println("Money in bank account: ", money)
	mutex.Unlock()
}
