package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// 1.
type byAge []user

func (bA byAge) Len() int {
	return len(bA)
}

func (bA byAge) Less(i, j int) bool {
	return bA[i].Age < bA[j].Age
}

func (bA byAge) Swap(i, j int) {
	bA[i], bA[j] = bA[j], bA[i]
}

type byLast []user

func (bL byLast) Len() int {
	return len(bL)
}

func (bL byLast) Less(i, j int) bool {
	return bL[i].Last < bL[j].Last
}

func (bL byLast) Swap(i, j int) {
	bL[i], bL[j] = bL[j], bL[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// 2.
	// normal
	for _, usr := range users {
		fmt.Printf("%v, %v %v\n", usr.Last, usr.First, usr.Age)
		for _, saying := range usr.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
	}

	fmt.Println("---------------------------------------------")

	// sorted by age
	sort.Sort(byAge(users))
	for _, usr := range users {
		fmt.Printf("%v, %v %v\n", usr.Last, usr.First, usr.Age)
		sort.Strings(usr.Sayings)
		for _, saying := range usr.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
	}

	fmt.Println("---------------------------------------------")

	// sorted by last name
	sort.Sort(byLast(users))
	for _, usr := range users {
		fmt.Printf("%v, %v %v\n", usr.Last, usr.First, usr.Age)
		sort.Strings(usr.Sayings)
		for _, saying := range usr.Sayings {
			fmt.Printf("\t%s\n", saying)
		}
	}
}
