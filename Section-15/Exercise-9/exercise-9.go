package main

import (
	"fmt"
)

func main() {
	// 1.
	res := mainFunc(func(x int) int {
		return x * x
	}, 26)

	fmt.Println(res)
}

func mainFunc(cb func(x int) int, n int) int {
	t := cb(n)
	return t + t
}
