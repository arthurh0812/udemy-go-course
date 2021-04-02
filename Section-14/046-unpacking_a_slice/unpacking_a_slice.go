package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 1, 8, 36, 45, 194, 22}

	// unpacking
	x := sum(xi...)

	fmt.Printf("Sum x: %v\n", x)

	// passing in nothing
	y := sum()

	fmt.Printf("Sum y: %v\n", y)
}

func sum(n ...int) int {
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	sum := 0
	for i, v := range n {
		sum += v
		fmt.Printf("i: %d\tv: %d\n", i, v)
	}
	return sum
}
