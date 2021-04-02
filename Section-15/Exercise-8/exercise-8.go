package main

import "fmt"

func main() {
	// 2.
	f1 := mainFunc("It's been a nice party")
	f1()

	f2 := mainFunc("The infections per day go higher and higher")
	f2()
	f2()
}

// 1.
func mainFunc(s string) func() {
	return func() {
		fmt.Printf("%s, as you've probably heard.\n", s)
	}
}
