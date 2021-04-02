package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	defer fmt.Println("Program exited.")

	c := make(chan int)

	wg.Add(2)

	// send
	go send(c)

	// receive
	go receive(c)

	wg.Wait()
}

func send(c chan<- int) {
	defer wg.Done()
	c <- 42
}

func receive(c <-chan int) {
	defer wg.Done()
	fmt.Println(<-c)
}
