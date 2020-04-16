package main

import "fmt"

// closure function for toggling boolean value
func toggleBool() func() bool {
	i := true
	return func() (ret bool) {
		ret = i
		if i {
			i = false
		} else {
			i = true
		}
		return
	}
}

// main function
func main() {
	// make a toggle closure
	tBool := toggleBool()
	fmt.Println(tBool())
	fmt.Println(tBool())
	fmt.Println(tBool())
	fmt.Println(tBool())
}
