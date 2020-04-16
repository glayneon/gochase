package main

import "fmt"

func main() {
	var ftemper, ctemper float32
	fmt.Print("Enter the temper in Fahrenheit: ")
	fmt.Scanf("%f", &ftemper)
	ctemper = ((ftemper - 32) * 5 / 9)
	fmt.Printf("The temperature in Celsius is %v.", ctemper)
}
