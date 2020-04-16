package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	x["key2"] = 20
	fmt.Println(x["key"])
}
