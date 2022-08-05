// count подсчитывает количество вхождений
// каждой строки в файле.
package main

import (
	"Go_Head_first/go/src/github.com/headfirstgo/datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
