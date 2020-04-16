package main

import (
	"fmt"
	"math/rand"
	"time"
)

// half
func half(x int) (int, bool) {
	const opr = 2
	var halves int
	halves = x / opr
	// conditionals
	if x%opr == 0 {
		return halves, true
	} else {
		return halves, false
	}
}

// main
func main() {
	rand.Seed(time.Now().UnixNano())
	tmp := rand.Intn(99)
	halves, even := half(tmp)
	fmt.Println(tmp, halves, even)
}
