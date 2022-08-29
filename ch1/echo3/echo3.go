package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Exercise 1.1
	fmt.Println("Name of the command that invokes this file:", os.Args[0])
	loopOverArguments()
	joinArguments()
}

func loopOverArguments() {
	// Exercise 1.2
	s, sep, start := "", "", time.Now()
	fmt.Println("Starting loopOverArguments: ", start)

	for i, arg := range os.Args[1:] {
		fmt.Printf("Index: %v, Value: %v", i, s)
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func joinArguments() {
	// Exercise 1.3 - benchmarking with time library
	start := time.Now()

	fmt.Println("Starting joinArguments: ", start)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
