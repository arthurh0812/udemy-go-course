package main

import (
	"fmt"
	"sync"
)

func main() {
	defer fmt.Println("Program exited.")

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range even {
			fanin <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range odd {
			fanin <- v
		}
	}()

	wg.Wait()
	close(fanin)
}
