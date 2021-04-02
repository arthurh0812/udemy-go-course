package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("something more")

	bar()
}

func foo() {
	fmt.Println("I'm in another block of code.")
}

func bar() {
	fmt.Println("and then we exited.")
}
