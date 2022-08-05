package main

import "fmt"

func main() {
	days := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for i, v := range days {
		fmt.Println(i, v)
	}
}
