package main

import (
	"fmt"
)

func test() {
	fmt.Println(c)
}

var c = 83
var d bool

func main() {
	// fmt.Println(a, b) - would both not work as there is no hoisting of variables inside non-package-level code

	// DECLARE a variable and ASSIGN a value (of a certain TYPE) at the same time

	// short declaration operator
	a := 5

	// var KEYWORD and assignment operator
	var b = 43

	fmt.Println(a, b)

	test()

	foo()

	fmt.Println(d)
}

func foo() {
	fmt.Println("again:", c)
	// fmt.Println("again:", a, b) - would not work as this block of code is outside the scope of a and b
}
