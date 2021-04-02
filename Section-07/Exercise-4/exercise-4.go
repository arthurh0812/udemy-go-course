package main

import "fmt"

func printBinDecHex(n int) {
	fmt.Printf("%d[dec]\t%b[bin]\t%#X[hex]\n", n, n, n)
}

func main() {
	// 1.
	n := 78

	// 2.
	printBinDecHex(n)

	// 3.
	m := n << 1

	// 4.
	printBinDecHex(m)
}
