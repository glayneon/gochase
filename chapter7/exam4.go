package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type recTangle struct {
	x1, x2, y1, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *recTangle) area() float64 {
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	a := Circle{1, 2, 3}
	b := recTangle{3, 7, 5, 12}
	fmt.Printf("Circle Area %f\n", a.area())
	fmt.Printf("Rectangle Area %f\n", b.area())
}
