package main

import "fmt"

func mapLookup(m map[string]int, key string) {
	if v, ok := m[key]; ok {
		fmt.Println(v)
		return
	}
	fmt.Printf("key does not exist: %q\n", key)
}

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 26,
	}

	mapLookup(m, "James")
	mapLookup(m, "Barnabas")
	mapLookup(m, "Miss Moneypenny")
}
