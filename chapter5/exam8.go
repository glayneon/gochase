package main

import "fmt"

func main() {
	var itmp int
	var ftmp float64
	var stmp string
	if itmp == 0 {
		fmt.Println("itmp", itmp)
	}
	if ftmp == 0 {
		fmt.Println("ftmp", ftmp)
	}
	if stmp == "" {
		fmt.Println("stmp", stmp)
	}
}
