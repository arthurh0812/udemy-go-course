package main

import (
	"fmt"
)

func main() {
	// 1.
	StateNamesOfUSA := make([]string, 50, 50)

	fmt.Println(len(StateNamesOfUSA))
	fmt.Println(cap(StateNamesOfUSA))

	for i := 0; i < len(StateNamesOfUSA); i++ {
		StateNamesOfUSA[i] = fmt.Sprintf("Position %d ", i)
	}

	fmt.Println(StateNamesOfUSA)
	fmt.Println(len(StateNamesOfUSA))
	fmt.Println(cap(StateNamesOfUSA))

	StateNamesOfUSA = append(StateNamesOfUSA, "Alibaba", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming")

	fmt.Println(StateNamesOfUSA)
	fmt.Println(len(StateNamesOfUSA))
	fmt.Println(cap(StateNamesOfUSA))

	// 2.
	for i := 0; i < len(StateNamesOfUSA)/2; i++ {
		fmt.Printf("%v %v\n", StateNamesOfUSA[i], StateNamesOfUSA[i+len(StateNamesOfUSA)/2])
	}

}
