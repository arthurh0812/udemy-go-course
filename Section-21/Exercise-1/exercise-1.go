package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 2.
var wg sync.WaitGroup

func main() {
	fmt.Println("I'm starting the main goroutine.")
	defer fmt.Println("I'm finishing the main goroutine.")

	fmt.Println("Goroutines(beginning):\t", runtime.NumGoroutine())

	wg.Add(2)

	// 1.
	go func() {
		runtime.Gosched()
		fmt.Println("Hi, there! I'm coming from goroutine 1.")
		wg.Done()
	}()

	go func() {
		runtime.Gosched()
		fmt.Println("Hi, there! I'm coming from goroutine 2.")
		wg.Done()
	}()

	fmt.Println("Goroutines(middle):\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Hi, there! I'm coming from the main goroutine.")
	fmt.Println("Goroutines(end):\t", runtime.NumGoroutine())
}
