package main

import "fmt"

// 1.
type person struct {
	first string
	last  string
	age   int
}

// 2.
func (p person) speak() {
	fmt.Printf("Hi, I'm %v %v and I am %d years old.\n", p.first, p.last, p.age)
}

func main() {
	// 3.
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	p.speak()

	p2.speak()
}
