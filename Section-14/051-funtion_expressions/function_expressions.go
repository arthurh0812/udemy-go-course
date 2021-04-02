package main

import "fmt"

func main() {
	// assigning a function to a variable
	f := func() {
		fmt.Println("My function expression")
	}

	g := func(x int) {
		fmt.Printf("%v\n", x)
	}

	// calling that variable
	f()

	g(4)
}
