package main

import "fmt"

func main() {
	// 1.
	arr := [5]int{1, 2, 3, 4, 5}

	// 2.
	for _, v := range arr {
		fmt.Println(v)
	}

	// 3.
	fmt.Printf("%T\n", arr)
}
