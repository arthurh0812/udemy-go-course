package main

import "fmt"

type person struct {
	first   string
	last    string
	favIces []string
}

func main() {

	p1 := person{
		first:   "James",
		last:    "Bond",
		favIces: []string{"Coconut", "Banana", "Martini"},
	}

	p2 := person{
		first:   "Miss",
		last:    "Monneypenny",
		favIces: []string{"Chocolat", "Lemon", "Smurf"},
	}

	// 1.
	mp := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	// 2.
	for _, v := range mp {
		printInfos(v)
	}
}

func printInfos(p person) {
	fmt.Println(p.first, p.first)
	for i, v := range p.favIces {
		fmt.Printf("\t%v: %v\n", i+1, v)
	}
	fmt.Println()
}
