/*
This is the case of struct type
*/

package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func (c *Circle) modify() {
	// func (c Circle) modify() {
	c.x = 5
}

// main
func main() {
	c := Circle{x: 0, y: 0, r: 5}
	c.modify()
	fmt.Println(c.x, c.y, c.r)
}
