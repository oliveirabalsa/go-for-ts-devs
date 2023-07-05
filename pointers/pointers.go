package main

import "fmt"

func main() {
	// Declare a variable and assign a value
	value := 42

	// Declare a pointer variable and assign the address of the variable
	pointer := &value

	// Access the value using the pointer
	fmt.Println("Value:", *pointer)

	// Modify the value using the pointer
	*pointer = 24
	fmt.Println("Modified value:", value)
}
