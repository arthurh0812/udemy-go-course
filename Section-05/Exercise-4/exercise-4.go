package main

import "fmt"

// 1.
type myOwnType int

// 2.
var x myOwnType

func main() {

	// 3.
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)
}
