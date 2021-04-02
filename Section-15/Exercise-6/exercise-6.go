package main

import "fmt"

func main() {
	// 1.
	func(x int) {
		fmt.Println(x * x * x)
	}(13)
}
