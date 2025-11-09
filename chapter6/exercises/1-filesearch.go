package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

func fileSearch(dir string, filename string, m *sync.Mutex, wg *sync.WaitGroup, paths *[]string) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			m.Lock()
			*paths = append(*paths, fpath)
			m.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, m, wg, paths)
		}
	}
	wg.Done()
}

func main() {
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	paths := []string{}
	go fileSearch(os.Args[1], os.Args[2], &m, &wg, &paths)
	wg.Wait()

	slices.Sort(paths)

	for _, fpath := range paths {
		fmt.Println(fpath)
	}
}
