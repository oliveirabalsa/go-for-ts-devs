package main

import "fmt"

func main() {
	// If statement
	num := 10
	if num > 0 {
		fmt.Println("Number is positive")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	// Switch statement
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good!")
	case "C":
		fmt.Println("Average")
	default:
		fmt.Println("Below Average")
	}
}
