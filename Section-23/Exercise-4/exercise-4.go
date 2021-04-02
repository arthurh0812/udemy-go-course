package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program exited.")

	q := make(chan int)
	c := gen(q)

	receive(c, q)
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Printf("from c:\t%v\n", v)
			}
		case v, ok := <-q:
			if ok && v == 1 {
				return
			}
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()

	return c
}
