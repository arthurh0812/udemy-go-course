package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program exited.")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(even, odd, quit)

	// receive
	receive(even, odd, quit)
}

func send(ec, oc chan<- int, qc chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			ec <- i
		} else {
			oc <- i
		}
	}
	close(ec)
	close(oc)
	qc <- true
	close(qc)
}

func receive(ec, oc <-chan int, qc <-chan bool) {
	for {
		select {
		case v, ok := <-ec:
			if ok {
				fmt.Println("from the eve channel:", v)
			}
		case v, ok := <-oc:
			if ok {
				fmt.Println("from the odd channel:", v)
			}
		case v, ok := <-qc:
			if ok {
				fmt.Println("from the quit channel:", v)
				return
			}
		}
	}
}
