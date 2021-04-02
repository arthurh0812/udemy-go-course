package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Outer: %d\n", i)
		fmt.Println()
		for j := 0; j < 5; j++ {
			fmt.Printf("\tInner: %d\n", j)
			// i += 2
		}
		fmt.Println()
	}
}
