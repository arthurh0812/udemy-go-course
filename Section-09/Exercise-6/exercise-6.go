package main

import (
	"fmt"
	"time"
)

const (
	birthYear = 2004
)

var (
	currentYear = time.Now().Year()
)

func main() {
	name := "James Bond"

	// 1.
	if (currentYear - birthYear) == 42 {
		fmt.Println("Lucky man!")
	}

	if birthYear < 2000 {
		fmt.Println("Really lucky man!")
	}

	if birthYear == 42 && name == "James Bond" {
		fmt.Println("You're fucking kidding me")
	}
}
