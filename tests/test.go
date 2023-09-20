package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)
	numbers = append(numbers, 4)
	fmt.Println(numbers)
}
