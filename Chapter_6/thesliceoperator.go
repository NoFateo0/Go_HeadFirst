package main

import "fmt"

// YOUR CODE HERE: Define a printLines function.
func printLines(lines []string) {
	for _, v := range lines {
		fmt.Println(v)
	}
}

func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}
	days := daysOfWeek[1:6]
	printLines(days)
	// YOUR CODE HERE: Get a slice containing just the
	// weekdays.
	// Pass that slice to printLines.
}
