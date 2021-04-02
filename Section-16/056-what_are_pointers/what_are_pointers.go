package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	// "&" gives you the address in memory
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	// "*" gives you the value which the following address is pointing to
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 44
	fmt.Println(a)
}
