package main

import (
	"fmt"
)

func main() {
	// 1.
	favs1 := []string{"Shaken, not stirred", "Martinis", "Women"}
	favs2 := []string{"James Bond", "Literature", "Computer Science"}
	favs3 := []string{"Being evil", "Ice cream", "Sunsets"}

	mp := map[string][]string{
		"bond_james":       favs1,
		"monneypenny_miss": favs2,
		"no_dr":            favs3,
	}

	// 2.
	for k, xs := range mp {
		fmt.Printf("Name: %q\n", k)
		for i, v := range xs {
			fmt.Printf("\ti: %d\tv: %q\n", i, v)
		}
	}
}
