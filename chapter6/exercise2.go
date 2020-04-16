package main

import (
	"fmt"
	"math/rand"
	"time"
)

// function that divide by 2 and return half value and remained number.
func returnHalf(x int) (int, string) {
	const divider = 2
	halfNum := x / divider
	if x%2 == 0 {
		return halfNum, "Even"
	} else {
		return halfNum, "Odd"
	}
}

// main function
func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(999)
	halfNum, oddEven := returnHalf(randNum)
	// print the result out
	fmt.Printf("Number: %d\nHalf: %d, %s\n", randNum, halfNum, oddEven)
}
