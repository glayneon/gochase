/*
This is a case of defer, recover and panic function.
*/

package main

import "fmt"

func first() {
	fmt.Println("First.")
}

func second() {
	fmt.Println("Second.")
}

// main
func main() {
	defer second()
	first()
}
