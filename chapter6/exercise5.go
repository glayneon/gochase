package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fibonacci function
func fibonacci(x int) int {
	if x == 0 {
		fmt.Printf("%d = ", x)
		return 0
	}
	fmt.Printf("%d + ", x)
	return x + fibonacci(x-1)
}

// main
func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)
	if randNum <= 3 {
		randNum += 3
	}
	sum := fibonacci(randNum)
	fmt.Printf("%d\n", sum)
	fmt.Printf("Number: %d\nFibonacci: %d\n", randNum, sum)
}
