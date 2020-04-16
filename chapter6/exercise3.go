package main

/*
This program is gonna find the greatest number in that slice
with variadic argument.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

// function that finding the greatest number in that slice with closure
func greatestNum(x ...int) (ret int) {
	maxNum := 0
	for _, i := range x {
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum
}

// function that generating random integers in slice
func genInts() []int {
	const maxNum = 12

	rand.Seed(time.Now().UnixNano())
	tmpSlice := make([]int, 0)

	for i, j := 0, rand.Intn(maxNum); i <= j; i++ {
		tmpSlice = append(tmpSlice, rand.Intn(100))
	}
	return tmpSlice
}

// main
func main() {
	tmpSlice := genInts()

	fmt.Print(tmpSlice)
	maxNum := greatestNum(tmpSlice...)
	fmt.Printf("The Greatest Number: %d\n", maxNum)
}
