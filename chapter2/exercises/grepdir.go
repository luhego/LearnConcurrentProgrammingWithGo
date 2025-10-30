package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	dir := args[2]

	fmt.Printf("Search term %s, dirname %s\n", term, dir)
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			go grep(term, path)
		}
		return nil

	})

	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

}
