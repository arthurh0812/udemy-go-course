package main

import (
	"fmt"
)

// 1.
type person struct {
	first string
	last  string
	age   int
	job   string
}

func main() {
	// 4.
	james := person{
		first: "James",
		last:  "Bond",
		age:   32,
		job:   "Secret Agent",
	}

	fmt.Println(james)

	// 5.
	changeMe(&james, options{
		first: "Thomas",
		age:   65,
		job:   "Banker",
	})

	fmt.Println(james)
}

type options struct {
	first string
	last  string
	age   int
	job   string
}

// 2.
func changeMe(p *person, opts options) {
	// 3.
	if v := opts.first; v != "" {
		(*p).first = v
	}
	if v := opts.last; v != "" {
		(*p).last = v
	}
	if v := opts.age; v != 0 {
		(*p).age = v
	}
	if v := opts.job; v != "" {
		(*p).job = v
	}
}
