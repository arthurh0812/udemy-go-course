package main

import (
	"fmt"
)

var y = 42
var z string = "Shaken, nor stirred"
var a string = `James said, "Shaken, 

not stirred"`

func main() {
	fmt.Println(y)
	// y = false
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// a = 45
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
