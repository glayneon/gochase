/*
This is the practice for the usage of interface in Go
*/

package main

import (
	"fmt"
	"math"
)

// Interface for Shape
type Shape interface {
	area() float64
	perimeter() float64
}

// structure for Circle
type Circle struct {
	radius float64
}

// method for calculating area of Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// method for calculating perimeter of Circle
func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

// structure for Rectangle
type recTangle struct {
	x, y float64
}

// method for calculating area of rectangle
func (r recTangle) area() float64 {
	return r.x * r.y
}

// method for calculating perimeter of rectangle
func (r recTangle) perimeter() float64 {
	return (r.x + r.y) * 2
}

// print Area and Perimeter of Shape
func printOut(s Shape) {
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("Area of this is %2f\n", s.area())
	fmt.Printf("Perimeter of this is %f\n\n", s.perimeter())
}

// main function
func main() {
	var s Shape = recTangle{5, 10}
	printOut(s)
	s = Circle{5}
	printOut(s)
}
