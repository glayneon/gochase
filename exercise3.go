package main

import "fmt"

var former, latter int32 = 32132, 42452

func main() {
	fmt.Printf("%v * %v = %v", former, latter, former*latter)
}
