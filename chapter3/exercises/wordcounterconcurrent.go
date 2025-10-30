package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type kv struct {
	k string
	v int
}

func countWords(url string, frequency map[string]int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	words := strings.Fields(string(body))

	for _, word := range words {
		frequency[word] += 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, frequency)
	}

	time.Sleep(10 * time.Second)

	frequencyList := make([]kv, 0, len(frequency))
	for w, f := range frequency {
		frequencyList = append(frequencyList, kv{k: w, v: f})
	}

	sort.Slice(frequencyList, func(i, j int) bool { return frequencyList[i].v > frequencyList[j].v })

	for i := range frequencyList {
		fmt.Printf("%s - %d\n", frequencyList[i].k, frequencyList[i].v)
	}
}
