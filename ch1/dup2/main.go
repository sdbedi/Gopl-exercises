package main

import (
	"bufio"
	"fmt"
	"os"
)

type Counter struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]*Counter)
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
		if n.Count > 1 {
			fmt.Printf("%d %v\n%s\n", n.Count, n.Filenames, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*Counter) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		_, ok := counts[input.Text()]
		if ok {
			counts[input.Text()].Count++
			counts[input.Text()].Filenames = append(counts[input.Text()].Filenames, f.Name())
		} else {
			counts[input.Text()] = new(Counter)
			counts[input.Text()].Count = 1
			counts[input.Text()].Filenames = append(counts[input.Text()].Filenames, f.Name())
		}

	}
	// NOTE: ignoring potential errors from input.Err()
}
