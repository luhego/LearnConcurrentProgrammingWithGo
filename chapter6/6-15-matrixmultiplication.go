package main

import (
	"fmt"
	"math/rand"
	"time"
)

const matrixSize = 2000

func matrixMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
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
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var A, B, C [matrixSize][matrixSize]int
	generateRandMatrix(&A)
	generateRandMatrix(&B)

	start := time.Now()
	matrixMultiply(&A, &B, &C)
	duration := time.Since(start)

	// fmt.Println("Matrix A:", A)
	// fmt.Println("Matrix B:", B)
	// fmt.Println("Result (A × B):", C)
	fmt.Println("Multiplication took:", duration)
}
