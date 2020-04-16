package main

import "fmt"

func add(xxx ...int) int {
	// declare variables
	var total int

	for _, value := range xxx {
		total += value
	}
	return total
}

func main() {
	fmt.Printf("Total is %d", add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
