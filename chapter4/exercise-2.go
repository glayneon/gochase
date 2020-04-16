package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// condtionals for Fizz Buzz
		switch tmp := i % 3; tmp {
		case 0:
			if i%5 == 0 {
				fmt.Printf("%d, FizzBuzz\n", i)
			} else {
				fmt.Printf("%d, Fizz\n", i)
			}
		default:
			if i%5 == 0 {
				fmt.Printf("%d, Buzz\n", i)
			} else {
				fmt.Printf("%d\n", i)
			}
		}
	}
}
