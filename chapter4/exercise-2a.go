package main

import "fmt"

// funciton main
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
			fmt.Print("\n")
		} else if i%5 == 0 {
			fmt.Print("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
