package main

import (
	"fmt"
	"runtime"
)

var x int8 = 127
var y float32
var z int64

func main() {
	x = 43
	// x = 42.893 -> would not work
	y = 42.45
	fmt.Println(x, y)
	fmt.Printf("Type of 'x': %T\n", x)
	fmt.Printf("Type of 'y': %T\n", y)

	// z = int(948923457896432) -> would not work, even though it has the same size

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
