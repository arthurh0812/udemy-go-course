package main

import "fmt"

const (
	by = 1 << (10 * iota)
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 3
	fmt.Printf("%d\t\t%b\n", y, y)

	a := 30 >> 2
	fmt.Printf("%d\t\t%b\n", a, a)

	fmt.Printf("%d\t\t%b\n", by, by)
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
