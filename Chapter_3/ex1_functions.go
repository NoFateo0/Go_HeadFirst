package main

import "fmt"

// YOUR CODE HERE: Write a scoreSummary function that will
// work with the calls below.
func scoreSummary(s string, a, b, c float64) (string, float64, float64, float64, float64) {
	sum := (a + b + c) / 3
	return s, a, b, c, sum
}

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("------------------------------------------------------")
	name, score1, score2, score3, average := scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	fmt.Printf("%10s | %8.1f | %8.1f | %8.1f | %8.2f\n", name, score1, score2, score3, average)
	name, score1, score2, score3, average = scoreSummary("Jessie", 79.3, 99.1, 82.5)
	fmt.Printf("%10s | %8.1f | %8.1f | %8.1f | %8.2f\n", name, score1, score2, score3, average)
	name, score1, score2, score3, average = scoreSummary("Lamar", 82.2, 95.4, 77.6)
	fmt.Printf("%10s | %8.1f | %8.1f | %8.1f | %8.2f\n", name, score1, score2, score3, average)

}
