package main

import (
	"fmt"
)

// 1.
type person struct {
	first string
	last  string
	age   int
}

// 2.
func (p *person) speak() {
	fmt.Printf("Hi, I'm %s %s.\n", p.first, p.last)
}

// 3.
type human interface {
	speak()
}

func main() {
	// 5.
	james := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	saySomething(&james)

	(&james).speak()

	// saySomething(james) => shows error

	james.speak() // => works though
}

// 4.
func saySomething(h human) {
	h.speak()
}
