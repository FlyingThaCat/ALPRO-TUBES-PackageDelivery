package main

import (
	"examples/math"
	"examples/output"
)

func main() {
	// Add two numbers
	sum := math.Add(5, 3)
	println("Sum:", sum)

	// Subtract two numbers
	difference := math.Subtract(5, 3)
	println("Difference:", difference)

	// Call the HelloWorld function
	output.HelloWorld()
	// Call the HelloUser function
	output.HelloUser()
}