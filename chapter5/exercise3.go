package main

import "fmt"

func main() {
	a := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var smallest int
	for _, value := range a {
		// conditional for initializing smallest variable
		if smallest == 0 {
			smallest = value
		} else {
			if value < smallest {
				smallest = value
			}
		}
	}
	fmt.Printf("The smallest number is %d", smallest)
}
