package main

import (
	"fmt"
)

var (
	a int
	c int
)

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	b = 13
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)

	// conversion works, as hotdog underlies type 'int'
	a = int(b)
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)

	c = 20

	// conversion works
	b = hotdog(c)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
}
