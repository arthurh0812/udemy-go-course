package main

import (
	"fmt"
)

func main() {
	favSport := "soccer"

	// 1.
	switch favSport {
	case "hockey":
		fmt.Println("Nice choice!")
	case "basketball":
		fmt.Println("Are you over 2m?")
	case "table tennis":
		fmt.Println("Reflexes are everything")
	case "soccer":
		fmt.Println("Great!")
	case "skiing":
		fmt.Println("Amazing")
	case "swimming":
		fmt.Println("Good for your body!")
	default:
		fmt.Println("I don't know that one...")
	}
}
