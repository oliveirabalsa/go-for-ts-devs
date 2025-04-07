package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	num := -8.7
	base := 2.0
	exp := 3.0

	// Absolute value
	abs := math.Abs(num)
	fmt.Println("Abs:", abs)

	// Rounding
	round := math.Round(num)
	fmt.Println("Round:", round)

	// Truncate
	trunc := math.Trunc(num)
	fmt.Println("Trunc:", trunc)

	// Ceiling
	ceil := math.Ceil(num)
	fmt.Println("Ceil:", ceil)

	// Floor
	floor := math.Floor(num)
	fmt.Println("Floor:", floor)

	// Square root
	sqrt := math.Sqrt(64)
	fmt.Println("Sqrt(64):", sqrt)

	// Power
	power := math.Pow(base, exp)
	fmt.Println("Pow(2, 3):", power)

	// Random number
	random := rand.Float64()
	fmt.Println("Random (0~1):", random)
}
