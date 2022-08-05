package main

import (
	"fmt"
	"log"
)

func perimeter(a, b float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("%f меньше нуля", a)
	} else if b < 0 {
		return 0, fmt.Errorf("%f меньше нуля", b)
	} else {
		return (a + b) * 2, nil
	}
}

func main() {
	var total float64

	result, err := perimeter(8.2, 10)
	if err != nil {
		log.Fatal(err)
	}
	total += result
	result, err = perimeter(5, 5.2)
	if err != nil {
		log.Fatal(err)
	}
	total += result
	result, err = perimeter(6.2, 4.5)
	if err != nil {
		log.Fatal(err)
	}
	total += result
	fmt.Println("You'll need", total, "meters of fencing")
}
