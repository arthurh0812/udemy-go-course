package main

import (
	"fmt"
)

var (
	s string
	i int
	b bool
	f float64
)

func main() {
	fmt.Printf("Printing the value of 's': %q.\n", s)
	fmt.Printf("Printing the type of 's': %T\n", s)
	s = "Bond, James Bond"
	fmt.Printf("Printing the value of 's': %q.\n", s)
	fmt.Printf("Printing the type of 's': %T\n", s)
	fmt.Printf("Printing the value of 'i': %v.\n", i)
	fmt.Printf("Printing the type of 'i': %T\n", i)
	fmt.Printf("Printing the value of 'f': %v.\n", f)
	fmt.Printf("Printing the type of 'f': %T\n", f)
	fmt.Printf("Printing the value of 'b': %t.\n", b)
	fmt.Printf("Printing the type of 'b': %T\n", b)
}
