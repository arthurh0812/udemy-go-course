package main

import "fmt"

func main() {
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// 1.
	slc1 := slc[:5]
	slc2 := slc[5:]
	slc3 := slc[2:7]
	slc4 := slc[1:6]

	fmt.Println(slc1)
	fmt.Println(slc2)
	fmt.Println(slc3)
	fmt.Println(slc4)
}
