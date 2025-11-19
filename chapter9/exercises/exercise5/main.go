package main

import (
	"chapter9/exercises/exercise1"
	"chapter9/exercises/exercise2"
	"chapter9/exercises/exercise3"
	"chapter9/exercises/exercise4"
)

func main() {
	quit := make(chan int)

	exercise4.Drain(quit,
		exercise3.Print(quit,
			exercise2.TakeUntil(func(s int) bool { return s <= 1000000 }, quit,
				exercise1.GenerateSquares(quit))))

	<-quit
}
