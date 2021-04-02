package main

import (
	"fmt"
)

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	x++
	fmt.Println(x)

	var y int
	fmt.Println(y)
	y++
	fmt.Println(y)

	{
		z := 0
		fmt.Println(z)
		z++
		fmt.Println(z)
	}

	// fmt.Println(z) doesn't work

	a := incrementor()
	b := incrementor()

	fmt.Printf("a: %v\n", a())
	fmt.Printf("a: %v\n", a())
	fmt.Printf("a: %v\n", a())
	fmt.Printf("a: %v\n", a())
	fmt.Printf("b: %v\n", b())
	fmt.Printf("b: %v\n", b())
}

func foo() {
	fmt.Println("Hello")
	x++
	// y++ doesn't work
	// z++ doesn't work
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
