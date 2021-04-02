package main

import (
	"fmt"
)

func main() {
	s := foo()
	fmt.Printf("%v\n", s)

	f := bar(67)
	fmt.Printf("Type: %T\n", f)
	fmt.Printf("%v\n", f())
}

func foo() string {
	return "Hello, World!"
}

// you can return a function from a function
// IMPORTANT: you then have to specify the exact type of the returned function as
// the return type(s) in your top-level function
func bar(n int) func() int {
	return func() int {
		return n
	}
}
