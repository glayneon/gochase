package main

import "fmt"

func main() {
	// declare variables, constants
	var feet float32
	const oneFeet = 0.3048

	// print the original value and the converted value
	fmt.Print("Enter the feet in order for converting to Meter: ")
	fmt.Scanf("%f", &feet)
	fmt.Printf("%f ft ==> %f", feet, feet*oneFeet)
}
