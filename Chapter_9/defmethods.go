package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func (r *rectangle) rectangleInfo() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func main() {
	var r rectangle
	r.length = 4.2
	r.width = 2.3
	// YOUR CODE HERE: Update this function call to a
	// method call.
	r.rectangleInfo()
}
