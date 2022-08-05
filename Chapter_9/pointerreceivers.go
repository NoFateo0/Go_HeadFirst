package main

func (r *rectangle) makeSquare() {
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

	var rectangle3 rectangle
	rectangle3.width = 10
	rectangle3.length = 100

	rectangle3.rectangleInfo()

	rectangle1.rectangleInfo()
	rectangle1.makeSquare()
	rectangle1.rectangleInfo()

	rectangle2.rectangleInfo()
	rectangle2.makeSquare()
	rectangle2.rectangleInfo()
}
