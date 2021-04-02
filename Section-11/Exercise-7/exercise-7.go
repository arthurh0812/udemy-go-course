package main

import (
	"fmt"
)

func main() {
	// 1.
	xs1 := []string{"James", "Bond", "Shaken, not stirred."}
	xs2 := []string{"Miss", "Monneypenny", "Hellooooo, James!"}

	xxs := [][]string{xs1, xs2}

	// 2.
	for i, xs := range xxs {
		fmt.Printf("Record: %d\n", i)
		for j, v := range xs {
			fmt.Printf("\ti: %d\tv: %q\n", j, v)
		}
	}
}
