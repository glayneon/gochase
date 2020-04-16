package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeOddGenerator() func() int {
	odd := 1
	return func() int {
		ret := odd
		odd += 2
		return ret
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nextOdd := makeOddGenerator()

	for i, j := 0, rand.Intn(9); i < j; i++ {
		fmt.Println(nextOdd())
	}
}
