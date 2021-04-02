package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello playground", 34, false)
	if err != nil {
		fmt.Println("something went wrong...")
	} else {
		fmt.Printf("%v", n)
	}
}
