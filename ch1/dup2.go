package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, arg, counts)
			f.Close()
		}

		for filename, count := range counts {
			for line, n := range count {
				if n > 1 {
					fmt.Printf("%s:\t%d\t%s\n", filename, n, line)
				}
			}
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if counts[filename] == nil {
			counts[filename] = map[string]int{}
		}

		counts[filename][input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
}
