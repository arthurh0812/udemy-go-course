package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1010)
	fmt.Println(x)

	x = append(x[:2], x[3:]...) // deletes x[2]
	fmt.Println(x)
}
