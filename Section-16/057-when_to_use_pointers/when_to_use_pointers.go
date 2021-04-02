package main

import "fmt"

func main() {
	n := 22

	fmt.Printf("Inside main: %v\n", n)
	foo(n)
	fmt.Printf("Inside main: %v\n", n)

	m := 0

	fmt.Printf("Inside main: %v\n", &m)
	fmt.Printf("Inside main: %v\n", m)
	bar(&m)
	fmt.Printf("Inside main: %v\n", &m)
	fmt.Printf("Inside main: %v\n", m)
}

func foo(y int) {
	fmt.Printf("Inside foo: %v\n", y)
	y = 43
	fmt.Printf("Inside foo: %v\n", y)
}

func bar(y *int) {
	fmt.Printf("Inside bar(pointer): %v\n", y)
	fmt.Printf("Inside bar(value): %v\n", *y)
	*y = 60
	fmt.Printf("Inside bar(pointer): %v\n", y)
	fmt.Printf("Inside bar(value): %v\n", *y)
}
