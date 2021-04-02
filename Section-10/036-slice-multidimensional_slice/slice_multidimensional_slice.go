package main

import "fmt"

func main() {
	// a multi-dimensional Slice is like a spreadsheet. You can have a slice of a
	// slice of some type.
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
