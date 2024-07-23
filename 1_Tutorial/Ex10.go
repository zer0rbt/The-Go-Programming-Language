// Fetchall fetches URLs in parallel and reports their times and sizes.\

// Exercise 1.10: Find a website that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do
// you get the same content each time? Modify fetchall to print its output to a file so it can be
// examined.

// Frodo: url to test: https://www.op.gg/champions
//
// Results
// 1st try: ~1.6s
// 2nd+ tries: ~0.8s, while size (and content) stays the same.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(strconv.Itoa(i)+".txt", url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(name string, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(name)
	if err != nil {
		ch <- fmt.Sprintf("while creating file %s: %v", name, err)
		return
	}

	nbytes, err := io.Copy(file, resp.Body)
	defer file.Close()
	// don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
