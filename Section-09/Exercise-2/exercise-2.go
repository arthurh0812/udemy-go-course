package main

import "fmt"

const (
	alphabetStart = 65
	alphabetEnd   = 90
)

var (
	times = 3
)

func main() {
	// 1.
	for i := alphabetStart; i <= alphabetEnd; i++ {
		fmt.Printf("%v(%#x)\n", i, i)
		for j := 0; j < times; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
