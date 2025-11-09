package main

import (
	"fmt"
	"math/rand"
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
