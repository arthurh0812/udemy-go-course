package main

import (
	"fmt"
)

type point struct {
	x float64
	y float64
}

func (p *point) calcY(fu func(x float64) float64) {
	p.y = fu(p.x)
	fmt.Println(p.y)
}

func main() {
	p := &point{
		x: 30,
	}

	p.calcY(f)

	fmt.Printf("%+v\n", p)
}

func f(x float64) float64 {
	return 1/120*(x*x) - x + 60
}
