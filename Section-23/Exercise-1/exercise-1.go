package main

import (
	"fmt"
)

func main() {
	solution1()
	solution2()
}

func solution1() {
	c := make(chan int)

	// 1.
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func solution2() {
	// 2.
	c := make(chan int, 1)

	c <- 37

	fmt.Printf("%v\n", <-c)
}
