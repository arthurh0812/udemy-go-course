package main

import "fmt"

func main() {
	// 1.
	slc := []int{2, 5, 8, 13, 24, 27, 92, 65, 76, 34}

	// 2.
	for _, v := range slc {
		fmt.Println(v)
	}

	// 3.
	fmt.Printf("%T\n", slc)
}
