package main

import "fmt"

func main() {
	// anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Susan",
		last:  "...",
		age:   64,
	}

	fmt.Println(p1)
}
