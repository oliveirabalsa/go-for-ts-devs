package main

import "fmt"

func main() {
	// Map creation and initialization
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35

	// Accessing values in a map
	aliceAge := ages["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	// Checking if a key exists in a map
	if age, ok := ages["Bob"]; ok {
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob's age not found")
	}

	// Deleting a key-value pair from a map
	delete(ages, "Charlie")
	fmt.Println("Charlie's age deleted")

	// Iterating over a map
	for name, age := range ages {
		fmt.Printf("%s's age: %d\n", name, age)
	}
}
