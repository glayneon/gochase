package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := Circle{rand.Float64() * 5, rand.Float64() * 5, rand.Float64() * 5}
	fmt.Println(c.area())
}
