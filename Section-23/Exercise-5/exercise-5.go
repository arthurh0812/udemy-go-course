package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
		close(c)
	}()

	v, ok := <-c
	fmt.Printf("v: %v\tok: %v\n", v, ok)

	v, ok = <-c
	fmt.Printf("v: %v\tok: %v\n", v, ok)
}
