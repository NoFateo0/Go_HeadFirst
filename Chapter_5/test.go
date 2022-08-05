package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[3] = "mi"
	fmt.Println(notes)

	var dates [3]time.Time

	dates[0] = time.Now()
	fmt.Println(dates[0])
}
