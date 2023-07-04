package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop (for loop without initialization and post statement)
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}
}
