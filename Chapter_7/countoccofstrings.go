package main

import "fmt"

func main() {
	// We'll count the number of times each letter occurs
	// within this slice.
	letters := []string{"b", "a", "c", "a", "b", "a",
		"a", "b", "c"}
	// YOUR CODE HERE: Process each element of "letters",
	// keeping track of how many times you've seen "a",
	// "b", or "c".
	// Finally, print out the number of times each letter
	// occurred.

	myMap := make(map[string]int)
	for _, v := range letters {
		myMap[v]++
	}
	fmt.Println(myMap)
}
