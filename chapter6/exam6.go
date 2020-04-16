package main

import "fmt"

func makeOddGenerator() func() int {
	x := 1
	return func() (ret int) {
		ret = x
		x += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
