package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Fname string   `json:"fname"`
	Lname string   `json:"lname"`
	Items []string `json:"items"`
}

func main() {
	s := `[
		{
			"fname": "James",
			"lname": "Bond",
			"items": [
				"Gin",
				"Suit",
				"Gun"
			]
		},
		{
			"fname": "Miss",
			"lname": "Moneypenny",
			"items": [
				"Dress",
				"Money",
				"Car"
			]
		}
	]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person

	unmarshal(bs, &people)

	fmt.Printf("%+v\n", people)

	for i, v := range people {
		fmt.Printf("i: %v\tv: %+v\n", i, v)
	}
}

func unmarshal(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(err)
	}
}
