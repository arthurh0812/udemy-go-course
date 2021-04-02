package main

import "fmt"

// 1.
const (
	a string = "James Smith"
	b        = 293
	c int32  = 186
	d        = "ok there"
)

func main() {
	// 2.
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%T\n", d)
}
