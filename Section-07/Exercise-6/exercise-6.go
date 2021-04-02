package main

import (
	"fmt"
)

// 1.
const (
	_         = iota
	nextYear1 = 2020 + iota
	nextYear2
	nextYear3
	nextYear4
)

func main() {
	sl := []int{nextYear1, nextYear2, nextYear3, nextYear4}

	// 2.
	for i, v := range sl {
		fmt.Printf("Next %v: %v\n", i+1, v)
	}
}
