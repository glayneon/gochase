package main

import (
	"fmt"
)

// swap function
func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

// main
func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Printf("x: %d\ty: %d\n", x, y)
}
