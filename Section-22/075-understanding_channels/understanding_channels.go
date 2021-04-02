package main

import (
	"fmt"
)

func main() {
	// doesNotWork()

	// usingGoroutines()

	// usingBuffers()

	unsuccesfulBuffers()
}

func doesNotWork() {
	c := make(chan int)

	// channel BLOCKS
	c <- 42

	fmt.Println(<-c)
}

func usingGoroutines() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func usingBuffers() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}

func unsuccesfulBuffers() {
	c := make(chan int, 1)

	c <- 42

	c <- 43

	fmt.Println(<-c)
}
