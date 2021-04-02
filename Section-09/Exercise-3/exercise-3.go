package main

import (
	"fmt"
	"time"
)

const (
	birthYear int16 = 2004
)

var (
	currentYear = int16(time.Now().Year())
)

func main() {
	var i int16 = birthYear

	// 1.
	for i <= currentYear {
		fmt.Println(i)
		i++
	}
}
