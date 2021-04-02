package main

import "fmt"

func main() {
	// 1.
	slc1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// 2.
	slc1 = append(slc1, 52)

	fmt.Println(slc1)

	// 3.
	slc1 = append(slc1, 53, 54, 55)

	fmt.Println(slc1)

	// 4.
	slc2 := []int{56, 57, 58, 59, 60}

	slc1 = append(slc1, slc2...)

	fmt.Println(slc1)
}
