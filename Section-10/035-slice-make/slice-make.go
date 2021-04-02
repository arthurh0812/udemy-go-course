package main

import "fmt"

func main() {
	x := make([]int, 12, 32)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x = append(x, 999, 80)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
