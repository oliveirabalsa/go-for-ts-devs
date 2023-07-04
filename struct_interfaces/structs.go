package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Struct initialization
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Person:", person)

	// Struct field access
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}
