package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 8, 42}

	fmt.Println(x)

	x = append(x, 7, 9, 10)

	fmt.Println(x)

	y := []int{234, 456, 678, 890}
	x = append(x, y...)

	fmt.Println(x)
}
