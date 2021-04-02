package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name string
	age  int
	job  string
}

func main() {
	f := foo()
	f(4)
	f(2)

	g := foo()
	g(2)

	f(1)
}

func foo() func(n int) {
	x := person{
		name: "Todd",
		age:  14,
		job:  "student",
	}
	var y int

	return func(n int) {
		for i := 1; i <= n; i++ {
			y++
			fmt.Printf("%v, ref: %v\n", x, unsafe.Pointer(&x))
		}
		fmt.Printf("%v\n", y)
	}
}
