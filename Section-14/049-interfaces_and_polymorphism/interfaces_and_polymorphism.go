package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return fmt.Sprintf("Hi, I'm person %s %s!", p.first, p.last)
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) lastPhrase() string {
	return fmt.Sprintf("Sorry, that...")
}

func (s secretAgent) speak() string {
	return fmt.Sprintf("Hi, I am secret agent %s %s!", s.first, s.last)
}

type monty struct {
	monopoly bool
	kopf     string
}

func (m monty) speak() string {
	return fmt.Sprintf("Hallo, ich bin ein Monopoly Man!")
}

type human interface {
	speak() string
}

func bar(h human) {
	switch h.(type) {
	// assertion
	case monty:
		fmt.Printf("I was passed into bar: %v\n", h.(monty).kopf)
	case secretAgent:
		fmt.Printf("I was passed into bar: %v\n", h.(secretAgent).first)
	default:
		fmt.Printf("I was passed into bar: %v\n", h.(person).first)
	}
}

type agent interface {
	speak() string
	lastPhrase() string
}

func foo(a agent) {
	fmt.Println(a)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.speak())
	fmt.Println(sa1.lastPhrase())

	bar(sa1)

	foo(sa1)

	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p1.speak())
	// fmt.Println(p1.lastPhrase()) does not exist

	bar(p1)

	// foo(p1) does not work

	m := monty{
		kopf:     "Kugel",
		monopoly: true,
	}

	bar(m)

	// foo(m) does not work

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
