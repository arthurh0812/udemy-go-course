package main

import "fmt"

func main() {
	// 1.
	a := (42 == 32)
	b := (42 <= 32)
	c := (42 >= 32)
	d := (42 != 32)
	e := (42 < 32)
	f := (42 > 32)

	// 2.
	fmt.Printf("%v %v %v %v %v %v", a, b, c, d, e, f)
}
