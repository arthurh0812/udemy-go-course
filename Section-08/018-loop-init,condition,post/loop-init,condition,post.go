package main

import "fmt"

func main() {
	// for <init>; <condition>; <post> {}

	for i := 0; i <= 100; i++ {
		fmt.Printf("Hello %v\n", i)
	}
}
