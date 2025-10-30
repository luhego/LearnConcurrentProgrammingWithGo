package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
* Adds 10 dollars 1 million times
 */
func stingy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money += 10
		runtime.Gosched()
	}
	fmt.Println("Stingy Done")
}

/*
* Substracts 10 dollars 1 million times
 */
func spendy(money *int) {
	for i := 0; i < 1000000; i++ {
		*money -= 10
		runtime.Gosched()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100

	go stingy(&money)
	go spendy(&money)

	time.Sleep(2 * time.Second)

	// We would expect money to 100 but due to a race condition, it is not
	println("Money in bank account: ", money)
}
