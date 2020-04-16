package main

import "fmt"

// closure function for generating Even number
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// main function
func main() {
	// make a Even generator
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
