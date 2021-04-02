package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licToKill bool
}

type mi7 struct {
	secretAgent
	id      string
	section string
}

func main() {
	person1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	agent1 := secretAgent{
		person:    person1,
		licToKill: true,
	}

	agent2 := secretAgent{
		person: person{
			first: "Johny",
			last:  "English",
			age:   43,
		},
		licToKill: false,
	}

	fmt.Println(agent1)

	fmt.Println(agent1.age, agent1.first, agent1.last, agent1.licToKill, agent1.person)

	fmt.Println(agent2)

	fmt.Println(agent2.age, agent2.first, agent2.last, agent2.licToKill, agent2.person)

	mp := mi7{
		secretAgent: secretAgent{
			person: person{
				first: "Dr.",
				last:  "Zhang",
				age:   35,
			},
			licToKill: true,
		},
		id:      "89v029vn56ls3",
		section: "Law",
	}

	fmt.Println(mp)

	fmt.Println(mp.age, mp.first, mp.last, mp.id, mp.licToKill, mp.person, mp.secretAgent, mp.section)
}
