package main

import (
	"fmt"
)

const (
	a = iota
	b
	c = 9
	d = iota
	e = 10
	f
)

const (
	g = iota
	h
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%T\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%T\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%T\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%T\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%T\n", h)
}
