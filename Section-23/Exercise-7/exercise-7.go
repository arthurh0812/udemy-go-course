package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

var num int = 10

func main() {
	defer log.Println("Program exited.")

	c := make(chan int)

	go spinGoroutines(c)

	receive(c)
}

func spinGoroutines(chnl chan<- int) {
	var wg sync.WaitGroup

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(c chan<- int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c <- rand.Intn(999)
			}
		}(chnl)
	}
	wg.Wait()
	close(chnl)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Printf("from channel:\t%v\n", v)
	}
}
