package main

import (
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.FprintF(os.StdErr, "dup3: %v\n", err)
			continue
		}
		for _, line := range string.Split(strings(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
