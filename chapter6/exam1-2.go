package main

import (
	"fmt"
	"math/rand"
	"time"
)

// average function with variadic arguments
func average(x ...float64) (float64, float64) {
	// variables for average and total
	var sum, avg float64
	// looping for total value
	for _, i := range x {
		sum += i
	}

	// return total values and average with len function
	avg = sum / float64(len(x))
	return sum, avg
}

// main
func main() {
	// variables for average and total value about tmpS slice
	var avg, total float64

	// slice for random integer under 99
	tmpS := make([]float64, 0)
	rand.Seed(time.Now().UnixNano())
	// looping for inserting value with append function
	for i, j := 0, rand.Intn(9); i <= j; i++ {
		tmpS = append(tmpS, float64(rand.Intn(99)))
	}

	fmt.Println(tmpS)
	total, avg = average(tmpS...)
	// Print total an average value out
	fmt.Printf("Total %f\nAverage %f\n", total, avg)

}
