package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func grep(term string, filename string) {
	fmt.Printf("Searching term %s in filename %s\n", term, filename)
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	if strings.Contains(string(data), term) {
		fmt.Printf("The filename %s contains the term %s\n", filename, term)
	}
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Not enough arguments")
	}

	term := args[1]
	for i := 2; i < len(args); i++ {
		go grep(term, args[i])
	}

	time.Sleep(2 * time.Second)
}
