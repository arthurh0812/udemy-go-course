package main

import (
	"fmt"
)

type person struct {
	name   string
	memory []string
}

func (p *person) tell(c chan<- string, message string) {
	c <- message
}

func (p *person) listen(c <-chan string) {
	for v := range c {
		p.memory = append(p.memory, v)
	}
}

func main() {
	defer fmt.Println("Program exited.")
	ch := make(chan int)

	go receiver(ch)

	ch <- 4

	// notWork(ch)

	test()
}

func notWork(ch chan int) {
	// BLOCKS and can only be stopped blocking by another (already running) goroutine
	ch <- 6

	go receiver(ch) // doesn't work as blocking this main goroutine in line 16 also
	// prevents this receiver goroutine func from being executed
}

func test() {
	ch := make(chan string)

	james := &person{
		name: "James Bond",
	}

	go james.listen(ch)

	hannah := &person{
		name: "Hannah Baker",
	}

	some := &person{
		name: "Some",
	}

	go some.listen(ch)

	go hannah.listen(ch)

	go hannah.tell(ch, "I'm really glad you did it.")

	go james.tell(ch, "I'm not so glad.")

	go james.tell(ch, "I should have done...")

	fmt.Printf("%+v\n", james)
	fmt.Printf("%+v\n", hannah)
	fmt.Printf("%+v\n", some)
}

func receiver(ch chan int) {
	fmt.Println("received from channel:", <-ch)
}
