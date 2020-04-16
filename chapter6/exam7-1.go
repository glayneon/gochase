/*
This case is one of the other cases for recursion programming
*/

package main

import "fmt"

// recursion function for calculating Factorial
func facto(x uint) uint {
	if x == 1 {
		return 1
	}
	return x * facto(x-1)
}

// main func
func main() {
	var tmpInt uint
	fmt.Print("Enter the number[1-9]: ")
	fmt.Scanf("%d", &tmpInt)

	// conditional for number is over 9 or less than 1
	if tmpInt <= 9 && tmpInt >= 1 {
		total := facto(tmpInt)
		fmt.Printf("Tthe number %d factorial is %d", tmpInt, total)
	} else {
		fmt.Printf("The number %d you've entered is wrong.", tmpInt)
	}
}
