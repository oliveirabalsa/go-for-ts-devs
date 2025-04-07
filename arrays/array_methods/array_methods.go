package main

import "fmt"

func main() {
	// Array declaration and initialization
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Append an element
	numbers = append(numbers, 6)
	fmt.Println("Appended Array:", numbers)

	// Copy an array/slice
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)
	fmt.Println("Copied Array:", copiedNumbers)

	// Slice an array/slice
	slicedNumbers := numbers[1:4] // [2, 3, 4]
	fmt.Println("Sliced Array:", slicedNumbers)

	// Modify the slice
	// Modifying a sliced array in Go, reflect on original array, because a slice is just a copy of array references;
	slicedNumbers[0] = 99
	fmt.Println("Modified Slice:", slicedNumbers)
	fmt.Println("Original Array:", numbers)

	// Flatten the nested slice (1-level only)
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Nested array:", nested)
	flat := flattenLevel(nested)
	fmt.Println("Flattened array:", flat)
}

func flattenLevel(input [][]int) []int {
	var flat []int
	for _, inner := range input {
		flat = append(flat, inner...)
	}
	return flat
}
