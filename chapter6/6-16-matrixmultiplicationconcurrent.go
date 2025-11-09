package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 2000

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int, row int, barrier *Barrier) {
	for {
		barrier.Wait()
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
		barrier.Wait()
	}
}

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func main() {
	start := time.Now()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var A, B, C [matrixSize][matrixSize]int
	barrier := NewBarrier(matrixSize + 1)
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&A, &B, &C, row, barrier)
	}

	generateRandMatrix(&A)
	generateRandMatrix(&B)
	barrier.Wait()
	barrier.Wait()

	duration := time.Since(start)

	fmt.Println("Multiplication took:", duration)
}

type Barrier struct {
	size      int        // Total number of goroutines required to trip the barrier
	waitCount int        // Number of goroutines currently waiting
	cond      *sync.Cond // Condition variable used to coordinate waiting and release
}

// NewBarrier creates and returns a new Barrier that will release
// goroutines once 'size' have called Wait.
func NewBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})
	return &Barrier{size, 0, condVar}
}

// Wait blocks the calling goroutine until the total number of waiting
// goroutines equals the barrier size. The last goroutine to arrive
// releases all waiting goroutines and resets the counter, allowing
// the barrier to be reused.
func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount++
	if b.waitCount < b.size {
		b.cond.Wait()
	} else {
		b.waitCount = 0
		b.cond.Broadcast()
	}
	b.cond.L.Unlock()
}

func workAndWait(name string, timeToWork int, barrier *Barrier) {
	start := time.Now()
	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}
