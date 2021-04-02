package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
		}()
		fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
