package main

import (
	"Go_Head_first/go/src/github.com/headfirstgo/datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("votes2.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// YOUR CODE HERE: Build a slice containing all the
	// key strings from the "counts" map.
	// Call the sort.Strings method on your slice.
	// Loop through the names in the sorted slice, and
	// print the name and the corresponding count from
	// the map.

	var mySlice []string
	for k := range counts {
		mySlice = append(mySlice, k)
	}

	sort.Strings(mySlice)
	for _, v := range mySlice {
		fmt.Printf("Votes for %s: %d\n", v, counts[v])
	}
}
