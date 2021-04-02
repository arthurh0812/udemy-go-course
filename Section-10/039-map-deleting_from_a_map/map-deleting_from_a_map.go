package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 26,
	}

	m["Todd"] = 33

	fmt.Println(m)

	delete(m, "James")

	fmt.Println(m)

	delete(m, "Barnabas")

	fmt.Println(m)
}
