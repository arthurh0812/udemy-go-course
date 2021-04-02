package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%v\n", x)
	x = 99
	fmt.Printf("%v\n", x)
	y := 100 + 24
	fmt.Printf("%v\n", y)
	z := "Bond, James"
	fmt.Printf("%q\n", z)
	b := true
	fmt.Printf("%t\n", b)
}
