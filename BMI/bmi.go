package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("BMI Berechnen")
	askHeight()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	height, _ := strconv.ParseFloat(scanner.Text(), 32)
	askWeight()
	scanner.Scan()
	weight, _ := strconv.ParseFloat(scanner.Text(), 32)

	bmi := weight / (height * height)
	fmt.Printf("Dein BMI ist: %v\n", bmi)

	if bmi <= 20 {
		fmt.Printf("You are underweight.")
	} else if bmi <= 25 {
		fmt.Printf("You are normal weight.")
	} else {
		fmt.Printf("You are overweight.")
	}

	scanner.Scan()
}

func askHeight() {
	x := "Please enter your height (in m):"
	fmt.Println(x)
}

func askWeight() {
	c := "Please enter your weight (in kg):"
	fmt.Println(c)
}
