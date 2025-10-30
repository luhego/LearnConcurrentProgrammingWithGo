package main

import (
	"fmt"
	"os"
	"time"
)

func outputFile(filename string) {
	fmt.Printf("Reading file %s\n", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func main() {
	args := os.Args

	for i := 1; i < len(args); i++ {
		go outputFile(args[i])
	}

	time.Sleep(2 * time.Second)
}
