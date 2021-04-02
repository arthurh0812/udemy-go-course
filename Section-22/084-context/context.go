package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------")

	ctx, cancel = context.WithDeadline(ctx, time.Now())

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------")

	ctx, cancel = context.WithTimeout(ctx, time.Second)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------")

	ctx = context.WithValue(ctx, "a", "b")

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Printf("cancel:\t%v\n", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)

	// mySleepAndTalk(ctx, 5*time.Second, "finished sleeping.")

	demo()
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func demo() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Printf("error check 1: %v\n", ctx.Err())
	fmt.Printf("num gortins 1: %v\n", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(200 * time.Millisecond)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Printf("error check 2: %v\n", ctx.Err())
	fmt.Printf("num gortins 2: %v\n", runtime.NumGoroutine())

	fmt.Println("about to cancel this context")
	cancel()
	fmt.Println("cancelled")

	time.Sleep(2 * time.Second)
	fmt.Printf("error check 3: %v\n", ctx.Err())
	fmt.Printf("num gortins 3: %v\n", runtime.NumGoroutine())

}
