package main

import "fmt"

func main() {
	x := sum(3, 4, 5, 6, 2, 8, 7, 6, 3)

	fmt.Printf("Sum: %v", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("i: %d\tv: %d\n", i, v)
	}

	return sum
}
