package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p person) speak() string {
	return fmt.Sprintf("Hello, I'm %s %s!\n", p.first, p.last)
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) lastPhrase() string {
	return fmt.Sprintf("I'm sorry, that...")
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(sa1.speak())
	fmt.Println(sa1.lastPhrase())

	fmt.Println(p1.speak())
	// fmt.Println(p1.lastPhrase()) isn't possible as the inheritance goes only in one
	// direction
}
