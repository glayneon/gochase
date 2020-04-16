/*
Factorial case
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// factorial function
func facto(x int) int {
	if x == 1 {
		return 1
	}
	return x * facto(x-1)
}

// main
func main() {
	const maxNum = 9
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(maxNum)
	factoValue := facto(randNum)
	fmt.Printf("Number: %d\nFactorial: %d\n", randNum, factoValue)
}
