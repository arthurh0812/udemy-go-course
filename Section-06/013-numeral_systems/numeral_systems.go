package main

import (
	"fmt"
)

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Printf("Type: %T\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%#b\n", n)
	fmt.Printf("%#x\n", n)

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	// fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Printf("%+q\n", sample)
}
