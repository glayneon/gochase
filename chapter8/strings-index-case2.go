package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello, World!"
	for _, i := range a {
		fmt.Printf("%v location in %v is %d\n", string(i), a, strings.Index(a, string(i)))
	}
}
