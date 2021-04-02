package main

import (
	"fmt"
)

// we create Values of a certain Type that are stored in Variables
// those Variables have Identifiers

var x int

type person struct {
	first string
	last  string
	age   int
}

type foo int

var y foo

// "Constant of a Kind"
const bar = 42

func main() {
	y = bar
	fmt.Printf("%T\n", y)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%v\n", bar)

	// variable of named type person assigned to anonymous type
	var p1 person = struct {
		first string
		last  string
		age   int
	}{
		first: "Ok",
		last:  "No",
		age:   34,
	}

	fmt.Println(p1)
}
