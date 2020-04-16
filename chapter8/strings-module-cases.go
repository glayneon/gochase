package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "test"
	b := []string{"test", "hello", "world", "Chase"}
	const text = "test message"
	// Contains (s, substr string) bool
	fmt.Println(strings.Contains(a, "es"))
	// => true

	// Count (s, sep string) int
	fmt.Println(strings.Count(a, "t"))
	// => 2

	// HasPrefix (s, prefix string) bool
	fmt.Println(strings.HasPrefix(a, "te"))
	// => true

	// HasSuffix (s, suffix string) bool
	fmt.Println(strings.HasSuffix(a, "st"))
	// => true

	// Index (s, index string) int
	fmt.Println(strings.Index(a, "s"))
	// => 2

	// Join (s []string, sep string) string
	fmt.Println(strings.Join(b, "="))
	// => test=hello=world=Chase

	// Repeat (s string, c count) string
	fmt.Println(strings.Repeat(text, 5))
	// => test messagetest messagetest messagetest messagetest message

	// Replace (s, old, new string, c count) string
	// if 'c' variable  value is -1 then it means change all of them which matched condition.
	fmt.Println(strings.Replace("aaaaaaaa", "a", "b", 3))
	// => bbbaaaaaa
	fmt.Println(strings.Replace("aaaaaaaa", "a", "b", -1))
	// => bbbbbbbbb
}
