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
	for {
		if !(i <= currentYear) {
			break
		}
		fmt.Println(i)
		i++
	}
}
