package main

import (
	"fmt"
)

// 1.
type vehicle struct {
	doors int
	color string
}

// 2.
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	// 3.
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "navy",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "orange",
		},
		luxury: false,
	}

	// 4.
	fmt.Println(t)
	fmt.Println(t.color)
	fmt.Println(s)
	fmt.Println(s.color)
}
