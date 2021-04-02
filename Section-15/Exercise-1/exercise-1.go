package main

import "fmt"

func main() {
	// 3.
	fi := foo(4)
	bi, bs := bar(13)
	fmt.Printf("Func foo: %v\n", fi)

	fmt.Printf("Func bar: %v\t%v\n", bi, bs)
}

// 1.
func foo(x int) int {
	return x * 34
}

// 2.
func bar(x int) (int, string) {
	msg := fmt.Sprintf("Hello, %v!\n", x)
	return x + 11, msg
}
