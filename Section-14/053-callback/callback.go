package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4, 5, 1, 5, 8, 7, 3, 13, 14}

	total1 := sum(numbers)
	fmt.Println(total1)

	total2 := even(sum, numbers)
	fmt.Println(total2)

	total3 := odd(sum, numbers)
	fmt.Println(total3)

	total4 := difference(numbers)
	fmt.Println(total4)

	total5 := even(difference, numbers)
	fmt.Println(total5)

	total6 := odd(difference, numbers)
	fmt.Println(total6)
}

func sum(xi []int) int {
	res := 0
	for _, v := range xi {
		res += v
	}
	return res
}

func difference(xi []int) int {
	res := xi[0]

	for _, v := range xi {
		res -= v
	}
	return res
}

// you can pass in a (callback) function as an argument which is then called inside
// the function
func even(cb func(x []int) int, y []int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}

	res := cb(xi)
	return res
}

func odd(cb func(x []int) int, y []int) int {
	var xi []int
	for _, v := range y {
		if v%2 != 0 {
			xi = append(xi, v)
		}
	}

	res := cb(xi)
	return res
}
