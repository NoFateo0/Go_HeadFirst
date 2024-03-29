package ordinals

import "fmt"

// Ordinal accepts an integer and returns a string with
// that integer's ordinal form.
func Ordinal(number int) string {
	lastDigit := number % 10
	isInTeens := (number%100)/10 == 1
	if isInTeens {
		return fmt.Sprintf("%dth", number)
	} else if lastDigit == 1 {
		return fmt.Sprintf("%dst", number)
	} else if lastDigit == 2 {
		return fmt.Sprintf("%dnd", number)
	} else {
		return fmt.Sprintf("%dth", number)
	}
}
