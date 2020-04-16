package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// declare stuff
	x := make([]int, 0)
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		x = append(x, i)
		fmt.Println(x)
	}
	// line delimeter
	fmt.Println("#############################")
	for i, _ := range x {
		// test
		tmp := rand.Intn(999)
		x[i] = tmp
		fmt.Println(x)
	}
}
