package main

import "fmt"

func main() {
	fmt.Printf("Nothing's up\n")
	showDefer()
	fmt.Printf("I am late...\n")
}

func showDefer() {
	defer func() {
		fmt.Printf("I am running inside the defer statement.\n")
	}()

	for i := 0; i <= 3; i++ {
		fmt.Printf("%d: I am inside the showDefer function.\n", i+1)
	}
}
