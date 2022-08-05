package main

import "fmt"

func main() {
	numbers := [9]int{1, 3, 10, 3, 1, 3, 3, 1, 10}
	var numArr [3]int

	for i := range numbers {
		numArr[i]++
	}

	for i, v := range numArr {
		fmt.Println(i, "occurred", v, "times")
	}
}
