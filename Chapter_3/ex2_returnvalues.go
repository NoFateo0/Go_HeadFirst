package main

import (
	"fmt"
)

// YOUR CODE HERE: Write a "perimeter" function.
func perimeter1(a, b float64) float64 {
	return (a + b) * 2
}

func main() {
	// YOUR CODE HERE: Call "perimeter" three times.
	// Add the three return values together, and store the
	// total in the "total" variable.
	var total float64
	total += perimeter1(8.2, 10)
	total += perimeter1(5, 5.2)
	total += perimeter1(6.2, 4.5)
	fmt.Println("You'll need", total, "meters of fencing")
}
