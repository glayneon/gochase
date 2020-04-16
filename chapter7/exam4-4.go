package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}

type recTangle struct {
	x, y float64
}

func (r recTangle) area() float64 {
	return r.x * r.y
}

func (r recTangle) perimeter() float64 {
	return (r.x + r.y) * 2
}

func printOut(s Shape) {
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("Area of s is %f\n", s.area())
	fmt.Printf("Perimeter of s is %f\n", s.perimeter())
}

func main() {
	var s Shape = recTangle{5, 7}
	printOut(s)
	s = Circle{5}
	printOut(s)
}
