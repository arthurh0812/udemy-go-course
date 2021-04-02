package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	defer fmt.Println("Program exited.")

	// even := make(chan int)
	// odd := make(chan int)
	// quit := make(chan bool)

	// wg.Add(2)

	// go send(even, odd, quit)

	// go receive(even, odd, quit)

	// wg.Wait()

	// commaOK()

	// selectReceive()

	selectSend()
}

// send channels
func send(even, odd chan<- int, quit chan<- bool) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	defer wg.Done()

	for {
		select {
		case v, ok := <-even:
			if ok {
				fmt.Println("even:", v)
			}
		case v, ok := <-odd:
			if ok {
				fmt.Println("odd:", v)
			}
		case _, ok := <-quit:
			if !ok {
				fmt.Println("quitting...")
				return
			}
		}
	}
}

func commaOK() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c

	fmt.Printf("%v %v\n", v, ok)

	v, ok = <-c

	fmt.Printf("%v %v\n", v, ok)
}

func selectReceive() {
	f := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 5000; i++ {
			f <- i
		}
	}()

	go func() {
		for i := 0; i < 5000; i++ {
			b <- i
		}
	}()

	foo(f, b)
}

func foo(f, b chan int) {
	defer fmt.Println("about to exit...")

	for i := 0; i < 10000; i++ {
		select {
		case v := <-f:
			fmt.Println("from f:", v)
		case v := <-b:
			fmt.Println("from b:", v)
		}
	}
}

func selectSend() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	bar(c, quit)
}

func bar(c, quit chan int) {
	x := 5
	for {
		select {
		case c <- x:
			x *= 2
		case <-quit:
			return
		}
	}
}
