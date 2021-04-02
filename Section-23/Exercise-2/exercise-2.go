package main

import (
	"fmt"
)

func main() {
	solution1()
	solution2()
}

func solution1() {
	// 1.
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func solution2() {
	// 2.
	cr := make(chan int)

	go func() {
		cr <- 37
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
