package main

import (
	"fmt"
	"os/exec"
	// "regexp"
	"strings"
	// "strconv"
	"bufio"
	// "math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		system(scanner)
	}
}

func system(s *bufio.Scanner) {
	s.Scan()
	expression := s.Text()

	if expression == "exit" {
		os.Exit(0)
	} else if expression == "clear" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		calculate(expression)
	}
}

func calculate(e string) {
	e = strings.ReplaceAll(e, " ", "")

	res := multiply(5, 6)
	fmt.Println(res)

	res = divide(10, 5, 3)
	fmt.Println(res)
}

func multiply(xf ...float64) float64 {
	var product float64 = 1
	for _, v := range xf {
		product *= v
	}
	return product
}

func divide(xd ...float64) float64 {
	var quotient float64 = xd[0]
	for _, v := range xd[1:] {
		quotient /= v
	}
	return quotient
}

func add(xs ...float64) float64 {
	var sum float64 = 0
	for _, v := range xs {
		sum += v
	}
	return sum
}

func subtract(xi ...float64) float64 {
	var difference float64 = xi[0]
	for _, v := range xi[1:] {
		difference -= v
	}
	return difference
}
