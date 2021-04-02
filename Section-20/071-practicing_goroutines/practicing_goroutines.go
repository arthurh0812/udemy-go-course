package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'g'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	// go numbers()
	// go alphabets()

	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("main terminated")

	channels()
}

func calcSquares(number int, sumop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	sumop <- sum
}

func calcCubes(number int, sumop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	sumop <- sum
}

func channels() {
	number := 589
	squarechan := make(chan int)
	cubechan := make(chan int)
	go calcSquares(number, squarechan)
	go calcCubes(number, cubechan)
	squares, cubes := <-squarechan, <-cubechan
	fmt.Println("Final output:", squares+cubes)

	// deadlock()

	// unidirectional()

	closingChannels()
}

func deadlock() {
	ch := make(chan int)
	ch <- 5
}

func unidirectional() {
	/* sendonlychnl is UNIdirectional in all functions */
	sendonlychnl := make(chan<- int)

	go sendData(sendonlychnl)
	// go readData(sendonlychnl)
	// go readAndSendData(sendonlychnl)

	/* readonlychnl is UNIdirectional in all functions */
	readonlychnl := make(<-chan int)

	// go sendData(readonlychnl)
	go readData(readonlychnl)
	// go readAndSendData(readonlychnl)

	/* chnl is BIdirectional in this function
	chnl is UNIdirectional (write) inside sendData()
	chnl is UNIdirectional (read) inside readData() */
	chnl := make(chan int)

	go sendData(chnl)        // => "channel conversion" (works only in one direction)
	go readData(chnl)        // => "channel conversion" (works only in one direction)
	go readAndSendData(chnl) // => normal channel
}

func sendData(sendchan chan<- int) {
	sendchan <- 10
}

func readData(readchan <-chan int) {
	num := <-readchan
	fmt.Println(num)
}

func readAndSendData(chnl chan int) {
	num := <-chnl
	fmt.Println(num)
	chnl <- 20
}

func closingChannels() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
