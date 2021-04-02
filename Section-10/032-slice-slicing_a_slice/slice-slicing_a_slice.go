package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 8, 42}
	fmt.Println(x)
	fmt.Println(x[:2])
	fmt.Println(x[2:])

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
