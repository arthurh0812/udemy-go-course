package main

import "fmt"

// 1.
type person struct {
	first   string
	last    string
	favIces []string
}

func main() {
	// 2.
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

	printInfos(p1)
	printInfos(p2)
}

func printInfos(p person) {
	fmt.Println(p.first, p.last)
	fmt.Printf("Favourite Ice Creams:\n")
	for i, v := range p.favIces {
		fmt.Printf("\t%v: %v\n", i+1, v)
	}
	fmt.Println()
}
