package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("This, that, and the other.")
	fmt.Printf("Type %T\nValues %v\n", words, words)
}
