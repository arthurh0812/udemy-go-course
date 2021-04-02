package main

import "fmt"

type point struct {
	x int32
	y int32
}

func main() {
	shape := struct {
		color   string
		corners []point
		edges   int
	}{
		color:   "red",
		corners: []point{point{x: 5, y: 7}, point{x: 5, y: 4}, point{x: 12, y: 7}, point{x: 12, y: 4}},
		edges:   4,
	}

	fmt.Println(shape.color, shape.edges)
	for i, p := range shape.corners {
		fmt.Printf("Point %d: (%d|%d)\n", i+1, p.x, p.y)
	}
}
