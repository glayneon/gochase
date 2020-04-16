package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello World!"
	for j, i := range string(a) {
		fmt.Printf("%d: %s => %s : %d\n", j, a, string(i), strings.Index(a, string(i)))
	}
}
