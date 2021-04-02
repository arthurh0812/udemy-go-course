package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// ColorGroup struct
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	// marshall
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%s\n", b)
	}

	// unmarshall
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll", "Order": "Dasyuromorphia"}
	]`)
	type animal struct {
		Name  string
		Order string
	}
	var animals []animal

	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("%+v\n", animals)
}
