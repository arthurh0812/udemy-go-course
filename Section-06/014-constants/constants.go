package main

import (
	"fmt"
)

const (
	a         = 42
	b float64 = 42
	c string  = "James Bond"
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", c)

	if a == b {
		fmt.Println("Nice")
	}
}
