package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

// byAge implements sort.Interface
type byAge []person

func (ba byAge) Len() int {
	return len(ba)
}

func (ba byAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (ba byAge) Less(i, j int) bool {
	return ba[i].age < ba[j].age
}

// byName implements sort.Interface
type byName []person

func (bn byName) Len() int {
	return len(bn)
}

func (bn byName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn byName) Less(i, j int) bool {
	return bn[i].first < bn[j].first
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Miss", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)
}
