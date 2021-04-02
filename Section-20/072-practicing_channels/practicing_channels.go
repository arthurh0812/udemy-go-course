package main

import (
	"fmt"
)

func digits(number int, dchn chan int) {
	for number != 0 {
		lastDigit := number % 10
		dchn <- lastDigit
		number /= 10
	}
	close(dchn)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	squarechn := make(chan int)
	cubechn := make(chan int)

	go calcSquares(number, squarechn)
	go calcCubes(number, cubechn)

	squares, cubes := <-squarechn, <-cubechn
	fmt.Println("Final output:", squares+cubes)
}
