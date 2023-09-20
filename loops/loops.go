package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("For Loop", i)
	}

	// While loop (for loop without initialization and post statement)
	j := 0
	for j < 5 {
		fmt.Println("While Loop no initialization", j)
		j++
	}

	// Infinite loop
	k := 0
	for {
		fmt.Println("Infinity with break", k)
		k++
		if k == 5 {
			break
		}
	}

	//For Range
	numbers := []int{3, 4, 5, 6, 7}
	for i, n := range numbers {
		fmt.Printf("For Range \n Index: %d, Number: %d\n", i, n)
	}

}
