package main

import (
	"fmt"
	"strings"
)

func main() {
	// String concatenation
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	// String length
	str := "Hello, World!"
	length := len(str)
	fmt.Println("Length:", length)

	// String splitting
	str = "Welcome to GoLang"
	split := strings.Split(str, " ")
	fmt.Println("Split result:", split)

	// String substring
	str = "Hello, World!"
	substr := str[0:5]
	fmt.Println("Substring:", substr)

	// String contains
	str = "Hello, World!"
	contains := strings.Contains(str, "World")
	fmt.Println("Contains:", contains)
}
