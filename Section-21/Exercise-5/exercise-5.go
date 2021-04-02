package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	incrementor int64
	waitgroup   sync.WaitGroup
)

const (
	numGoroutines = 50
)

func main() {
	defer fmt.Println("Program exited.")

	waitgroup.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go increment()
	}

	waitgroup.Wait()
}

func increment() {
	defer waitgroup.Done()

	atomic.AddInt64(&incrementor, 1)
	runtime.Gosched()
	fmt.Println("Incrementor:", atomic.LoadInt64(&incrementor))
}
