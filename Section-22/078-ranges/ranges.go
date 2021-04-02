package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program exited.")

	c := make(chan int)

	// send
	go send(c)

	// range loops stop reading from a channel when the channel is closed
	for v := range c {
		fmt.Println(v)
	}
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
