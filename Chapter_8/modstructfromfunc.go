package main

import "fmt"

type rectangle struct {
	width  float64
	length float64
}

func rectangleInfo(r rectangle) {
	fmt.Println("Width: ", r.width)
	fmt.Println("Length: ", r.length)
}

func makeSquare(r *rectangle) {
	if r.width < r.length {
		r.length = r.width
	} else {
		r.width = r.length
	}
}
func main() {
	var rectangle1 rectangle
	rectangle1.width = 100
	rectangle1.length = 50

	var rectangle2 rectangle
	rectangle2.width = 20
	rectangle2.length = 40

	makeSquare(&rectangle1)
	rectangleInfo(rectangle1)

	makeSquare(&rectangle2)
	rectangleInfo(rectangle2)
}
