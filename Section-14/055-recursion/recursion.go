package main

import "fmt"

const n int = 6

func main() {
	// solution using recursion
	a := factorial(n)
	fmt.Println(a)

	// solution using loop
	b := loopFactorial(n)
	fmt.Println(b)
}

// a recursion is a function that calls itself (kind of infinite except you specify
// condition when to no longer call itself)
func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	res := 1
	for ; n > 0; n-- {
		res *= n
	}
	return res
}
