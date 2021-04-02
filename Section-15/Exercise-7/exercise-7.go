package main

import "fmt"

func main() {
	// 1.
	f := func(x float64, y int) []float64 {
		var res []float64
		for i := 1; i <= y; i++ {
			res = append(res, float64(i)*x)
		}
		return res
	}

	fmt.Printf("%T\n", f)

	res := f(13.5, 10)

	fmt.Println(res)
}
