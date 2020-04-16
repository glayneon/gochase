package main

import "fmt"

// main func
func main() {
	// declare the array x
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	var total float64 = 0
	for _, value := range x {
		total += value
	}

	// print the average of total
	fmt.Println(total / float64(len(x)))
}
