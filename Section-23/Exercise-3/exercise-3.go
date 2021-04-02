package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program exited.")

	c := gen()
	receive(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Printf("from c:\t%v\n", v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
