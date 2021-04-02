package main

import "fmt"

// 1.
var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
)

func main() {
	// 2.
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
