/*
This is a case of recursive function usage in Golang.

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// recursion function declaration
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// main
func main() {
	rand.Seed(time.Now().UnixNano())
	tmp := rand.Intn(10)

	// print the result of the factorial
	fmt.Printf("Number: %d\nFactorial: %d\n", tmp, factorial(tmp))
}
