package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 13
	fmt.Printf("%#x\n%b\n%x\n", y, y, y)

	var i interface{} = 23

	fmt.Printf("%v\n", i)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
}
