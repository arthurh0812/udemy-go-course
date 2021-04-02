package main

import (
	"fmt"
)

func divide(n int, m int) float64 {
	return float64(n) / float64(m)
}

func main() {
	x := 83
	y := 49

	a := divide(x, y)

	fmt.Println(a)

	x = -1
	for {
		x++

		if x >= 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
}
