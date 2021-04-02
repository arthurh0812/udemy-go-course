package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

// method set of type circle is:
// type methodSetCircle interface {}

// method set of type pointer to a circle is:
// type methodSetPointerCircle interface {
// 	 area() float64
// }

func main() {
	c1 := circle{
		radius: 7.5,
	}

	c2 := &circle{
		radius: 10,
	}

	// non-pointer value c has access to the area method (although the receiver is of
	// type *circle)
	fmt.Println(c1.area())

	fmt.Println(c2.area())

	// non-pointer value c does not implement the shape interface though
	// info(c1) => error

	info(c2) // => works

	// => the method set only determines the INTREFACES that this Type implements (not the functions that is can access as it cann access all that are declared with receiver of type T or *T)
}

func info(s shape) {
	fmt.Println("Area:", s.area())
}
