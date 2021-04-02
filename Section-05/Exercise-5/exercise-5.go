package main

import "fmt"

// 1.
type myOwnType int

var (
	x myOwnType
	y int
)

func main() {
	// 2.
	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
