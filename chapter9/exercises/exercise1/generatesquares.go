package exercise1

import "fmt"

func GenerateSquares(quit <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		defer close(squares)
		for i := 1; ; i++ {
			select {
			case squares <- i * i:
			case <-quit:
				fmt.Println("Quitting squares generation")
				return
			}
		}

	}()

	return squares
}
