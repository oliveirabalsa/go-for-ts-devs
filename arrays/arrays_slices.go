package main

import "fmt"

func main() {
	// Array declaration and initialization
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println("Array:", numbers)

	// Array literal
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Array:", names)

	// Array length
	fmt.Println("Array Length:", len(numbers))

	// Iterating over an array
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Slicing an array
	slicedNumbers := numbers[1:4] // [2, 3, 4]
	fmt.Println("Sliced Array:", slicedNumbers)

	// Modifying the slice
	slicedNumbers[0] = 99
	fmt.Println("Modified Slice:", slicedNumbers)
	fmt.Println("Original Array:", numbers)
}
