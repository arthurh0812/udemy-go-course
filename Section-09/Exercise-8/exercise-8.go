package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 == 5:
		fmt.Println("ok")
	case 4*7 == 42:
		fmt.Println("NOT ok")
	case true:
		fmt.Println("very nice")
	case false:
		fmt.Println("this is a switch")
	default:
		fmt.Println("this gets printed when none of the case expressions matched")
	}
}
