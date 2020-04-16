package main

import (
	"fmt"
	"math"
)

type Myfloat float64

func (m Myfloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	} else {
		return float64(m)
	}
}

func main() {
	m := Myfloat(-math.Sqrt2)
	m2 := Myfloat(-2.124)
	fmt.Println(m.Abs())
	fmt.Println(m2.Abs())
}
