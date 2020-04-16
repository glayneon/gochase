/*
This is a case of defer, recover and panic function.
*/

package main

import "fmt"

// main
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
