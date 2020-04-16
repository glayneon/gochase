package main

import (
	"fmt"
	"math/rand"
	"time"
)

// average funciton
func average(xs []float64) (float64, float64) {
	// declare variables
	var total, average float64

	// for loop for finding the average value
	for _, j := range xs {
		total += j
	}
	average = total / float64(len(xs))

	// return thoes variables
	return total, average
}

// main function
func main() {
	tempSlice := make([]float64, 0)
	// make a seed
	rand.Seed(time.Now().UnixNano())
	// make a length how long this have length
	for i, j := 0, rand.Intn(10); i <= j; i++ {
		tmp := float64(rand.Intn(99))
		tempSlice = append(tempSlice, tmp)
	}

	// print the current values of that slice
	fmt.Println(tempSlice)

	// print total value and average of total value
	total, avg := average(tempSlice)
	fmt.Printf("Total: %.1f\nAverage: %.1f\n", total, avg)
}
