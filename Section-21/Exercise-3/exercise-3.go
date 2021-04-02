package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int

var waitgroup sync.WaitGroup

const (
	gs = 50
)

func main() {
	defer fmt.Println("Program exited.")

	waitgroup.Add(gs)

	for i := 0; i < gs; i++ {
		go increment()
	}

	waitgroup.Wait()

	fmt.Printf("Incrementor: %v\n", incrementor)
}

func increment() {
	defer waitgroup.Done()
	v := incrementor
	runtime.Gosched()
	v++
	incrementor = v
}
