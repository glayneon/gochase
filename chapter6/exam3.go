package main

import "fmt"

// add function with variadic parameters
func add(xxx ...int) int {
	var total int
	for _, i := range xxx {
		total += i
	}
	return total
}

// main
func main() {
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}
