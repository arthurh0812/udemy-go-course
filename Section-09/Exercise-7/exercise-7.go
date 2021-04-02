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
	// 1.
	if (currentYear - birthYear) > 80 {
		fmt.Println("Youre having a pretty long life!")
	} else if (currentYear - birthYear) > 60 {
		fmt.Println("Youve still got a good amount of years ahead of you...")
	} else if (currentYear - birthYear) > 40 {
		fmt.Println("You are in your golden ages!")
	} else if (currentYear - birthYear) > 20 {
		fmt.Println("Young man/woman...")
	} else {
		fmt.Println("Are you still in the kindergarden? Or not even born yet?")
	}
}
