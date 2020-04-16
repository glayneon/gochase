package main

import (
	"fmt"
	"math"
)

const multiple = 8

func main() {
	answer := math.Pow(2, multiple) - 1
	fmt.Printf("11111111 in base 2 digit is converting to %v in base 10 digit", answer)
}
