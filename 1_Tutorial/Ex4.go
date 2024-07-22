// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type substring struct {
	count int
	fname map[string]int
}

func (p *substring) Add(file string) {
	if p.fname == nil {
		p.fname = make(map[string]int)
	}
	p.fname[file]++
	p.count++
}

func main() {
	counts := make(map[string]substring)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t", n.count, line)
			for fname := range n.fname {
				fmt.Printf("%s ", fname)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]substring) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		sub := counts[line]
		sub.Add(f.Name())
		counts[line] = sub
	}
}
