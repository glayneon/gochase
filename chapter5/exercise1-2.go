package main

import "fmt"

func main() {
	a := make([]int, 3, 9)
	fmt.Printf("Length: %d\nContent: %v", len(a), a)
	// exercise 2
	b := [6]string{"a", "b", "c", "d", "e", "f"}
	c := b[2:5]
	fmt.Print(c)
}
