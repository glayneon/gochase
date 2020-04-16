package main

import "fmt"

type aaa struct {
	x, y int
}

func (a aaa) square() {
	a.x = 5
}

func (a *aaa) square2() {
	a.x = 5
}

func main() {
	a := aaa{1, 2}
	fmt.Println("1. ", a)
	a.square()
	fmt.Println("2 ", a)
	a.square2()
	fmt.Println("3. ", a)
}
