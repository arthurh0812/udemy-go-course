package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
	job   string
}

func (p person) greet() string {
	return fmt.Sprintf("Hi, I'm %s %s, %s, and I am %d years old.",
		p.first,
		p.last,
		p.job,
		p.age,
	)
}

func (p *person) sayHello() {
	fmt.Printf("Hello, there, I'm %s!\n", p.first)
}

type human interface {
	greet() string
}

type pointer interface {
	sayHello()
}

func main() {
	p := person{
		first: "Ani",
		last:  "Achola",
		job:   "student",
		age:   17,
	}

	greeting1 := (p).greet()
	greeting2 := (&p).greet()
	fmt.Println(greeting1)
	fmt.Println(greeting2)

	p.sayHello()
	(&p).sayHello()

	info(p)

	info(&p)

	// isPointer(p) shows that p is Not of type pointer (interface)

	// that means pointers to a certain type have the method set of that type, but not
	// the other way round (the type doesn't have the methods of pointers to that
	// type included in its method set)

	isPointer(&p)
}

func info(h human) {
	fmt.Println(h)
}

func isPointer(p pointer) {
	fmt.Println(p)
}
