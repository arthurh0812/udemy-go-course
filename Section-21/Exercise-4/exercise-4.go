package main

import (
	"fmt"
	"sync"
)

var incrementor int

var waitgroup sync.WaitGroup

// 1.
var mtx sync.Mutex

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
	mtx.Lock()

	v := incrementor
	v++
	incrementor = v

	mtx.Unlock()
}
