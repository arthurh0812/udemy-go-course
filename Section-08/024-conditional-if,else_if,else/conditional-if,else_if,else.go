package main

import "fmt"

func main() {
	x := 40
	if x == 40 {
		fmt.Println("the value of x was 40")
	} else if x == 41 {
		fmt.Println("the value of x was not 40, but 41")
	} else {
		fmt.Println("the value of x was neither 40, nor 41")
	}
	fmt.Println(x)

	y := 458739823 >> 21
	if y < 50 {
		fmt.Println("the value of y was small (less than 50)")
	} else if y < 100 {
		fmt.Println("the value of y was rather small (not less than 50, but less than 100")
	} else if y < 150 {
		fmt.Println("the value of y was rather great (neither less than 50, nor less than 100, but less than 150")
	} else if y < 200 {
		fmt.Println("the value of y was great (neither less than 50, nor less than 100, nor less than 150, but less than 200")
	} else {
		fmt.Println("the value of y was really huge (neither less than 50, nor less than 100, nor less than 150, nor less than 200)")
	}
	fmt.Println(y)
}
