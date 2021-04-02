package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Monneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	fmt.Println(p2.age)
}
