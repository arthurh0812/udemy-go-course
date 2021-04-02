package main

import "fmt"

func main() {
	xi := []int{3, 5, 2, 7, 4, 3, 2, 6, 1, 4, 3}

	// 2.
	s1 := foo(xi...)
	fmt.Println(s1)

	// 4.
	s2 := bar(xi)
	fmt.Println(s2)
}

// 1.
func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// 3.
func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
