package main

import (
	"fmt"
	"math/rand"
)

func main() {
	defer fmt.Println("Program exited.")

	c := make(chan int)

	go send(c)

	receive(c)
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- rand.Intn(1000)
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Printf("from c:\t%v\n", v)
	}
}
