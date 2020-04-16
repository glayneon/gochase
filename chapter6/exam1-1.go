package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this's gonna return the set of values with total sum and average.
func average(args ...float64) (float64, float64) {
	var total float64
	// calculate total value
	for _, i := range args {
		total += i
	}

	return total, total / float64(len(args))
}

// main
func main() {
	rand.Seed(time.Now().UnixNano())
	tmpSlice := make([]float64, 0)
	for i, j := 0, rand.Intn(9); i < j; i++ {
		tmpSlice = append(tmpSlice, float64(rand.Intn(999)))
	}
	fmt.Println(tmpSlice)
	total, avg := average(tmpSlice...)
	fmt.Printf("Total: %.1f\nAverage: %.1f\n", total, avg)
}
