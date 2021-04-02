package main

import (
	"fmt"
	"math"
)

// 1.
type square struct {
	width float32
}

// 3.
func (s square) area() float32 {
	return s.width * s.width
}

func (s square) scope() float32 {
	return 4 * s.width
}

// 2.
type circle struct {
	radius float32
}

// 3.
func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) scope() float32 {
	return math.Pi * 2 * c.radius
}

// 4.
type shape interface {
	area() float32
	scope() float32
}

func main() {
	// 6.
	sq := square{
		width: 5.4,
	}

	ci := circle{
		radius: 4.5,
	}

	ci2 := circle{
		radius: 43,
	}

	// 7.
	infos := info(sq)
	fmt.Printf("area: %v\tscope: %v\n", infos.area, infos.scope)

	infos = info(ci)
	fmt.Printf("area: %v\tscope: %v\n", infos.area, infos.scope)

	infos = info(ci2)
	fmt.Printf("area: %v\tscope: %v\n", infos.area, infos.scope)
}

type infoStruct struct {
	area  float32
	scope float32
}

// 5.
func info(s shape) infoStruct {
	return infoStruct{
		area:  s.area(),
		scope: s.scope(),
	}
}
