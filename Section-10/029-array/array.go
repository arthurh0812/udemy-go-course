package main

import (
	"fmt"
)

func main() {
	var x [10]int
	fmt.Println(x)
	x[0] = 4
	x[3] = 203
	x[len(x)-1] = 56
	fmt.Println(x)
	fmt.Println(len(x))
}
