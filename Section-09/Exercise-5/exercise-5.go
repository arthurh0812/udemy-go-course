package main

import (
	"fmt"
)

var (
	beginning int = 10
	end       int = 100
	divisor   int = 4
)

func main() {
	// 1.
	for i := beginning; i <= end; i++ {
		mod := i % divisor
		fmt.Printf("For %d: remainder %d\n", i, mod)
	}
}
