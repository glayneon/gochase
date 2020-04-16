package main

import "fmt"

func main() {
	var answer int

	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &answer)

	switch answer {
	case 0:
		fmt.Print("Zero")
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	case 4:
		fmt.Print("Four")
	default:
		fmt.Print("Unknown Number..")
	}
}
