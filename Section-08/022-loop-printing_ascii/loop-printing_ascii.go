package main

import "fmt"

func main() {
	fmt.Printf("Bin\tDec\tHex\tUnicode\n")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%b\t%d\t%x\t%#U\n", i, i, i, i)
	}
}
