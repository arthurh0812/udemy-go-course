package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program exited.")
	invalid1()
	invalid2()

	cb := make(chan int)   // bi
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("%T\n", cb)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// "from specific to general" does not work
	// cb = cr
	// cb = cs

	// "from general to specific" does work (in this case: "Channel Conversion")
	cs = cb
	cr = cb
}

func invalid1() {
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// invalid operation
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

func invalid2() {
	c := make(<-chan int, 2)

	// invalid operation
	// c <- 42
	// c <- 43

	// fmt.Println(<-c)
	// fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

func goroutines() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
