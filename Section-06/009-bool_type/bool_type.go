package main

import "fmt"

var x bool

func main() {
	fmt.Printf("%v\n", x)
	x = true
	fmt.Printf("%v\n", x)

	a := 0
	b := 3
	c := a == b
	d := a < b
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}
