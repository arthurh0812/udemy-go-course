package main

import (
	"fmt"
)

func main() {
	s := `Hello "Playground`

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	s = `Hello "Go`

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%d :\t'%c'\n", bs[i], bs[i])
	}

	fmt.Println()

	for i, v := range s {
		fmt.Printf("%d. hex:'%X' and uni:'%d'\n", i, v, v)
	}
}
