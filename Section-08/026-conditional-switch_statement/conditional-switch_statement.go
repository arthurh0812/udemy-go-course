package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("1")
	case (2+2 == 3):
		fmt.Println("2")
	case (3 == 3):
		fmt.Println("3")
		fallthrough
	case (4 == 4):
		fmt.Println("4")
		fallthrough
	case (3 == 5):
		fmt.Println("5")
	case (9 == 7):
		fmt.Println("6")
	case (15 == 15):
		fmt.Println("7")
	default:
		fmt.Println("default")
	}

	n := "Bond"

	switch n {
	case "Moneypenny", "bond", "No":
		fmt.Println("miss money or bond or no")
	case "Bond":
		fmt.Println("bond, james")
	case "Q":
		fmt.Println("q")
	default:
		fmt.Println("this is default")
	}
}
