package main

import "fmt"

func main() {
	// not anonymous
	foo()

	// anonymous, self-executing functions
	func() {
		fmt.Printf("first anonymous func ran!\n")
	}()

	func(p string) {
		fmt.Printf("second anonymous func ran! %s\n", p)
	}("Wohoo!")

	func(x int, a ...interface{}) {
		fmt.Printf("third anonymous func ran! %d. %v\n", x, a)
	}(42, struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	})
}

func foo() {
	fmt.Printf("foo ran.\n")
}
