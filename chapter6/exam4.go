/*
It's a case of usage for Closure in Golang.
*/

package main

import "fmt"

func main() {
	// declare variable x
	x := 0

	// closure
	incr := func() int {
		x++
		return x
	}

	// print out
	fmt.Println(incr())
	fmt.Println(incr())
}
