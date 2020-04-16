package main

import "fmt"

const greeting = "Hello World"

func main() {
	fmt.Println(len(greeting))
	fmt.Println(greeting[1])
	// it's gonna print "Hello World Chase" in three lines
	fmt.Println("Hello\nWorld\nChase!")
	fmt.Println("Hello" + "World")
}
