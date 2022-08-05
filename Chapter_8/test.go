package main

import (
	"Go_Head_first/go/src/github.com/headfirstgo/magazine"
	"fmt"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	//subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	//subscriber.State = "NE"
	//subscriber.PostalCode = "68111"
	//fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	//fmt.Println("State:", subscriber.State)
	//fmt.Println("Postal Code:", subscriber.PostalCode)
	fmt.Println(subscriber)

	employee := magazine.Employee{Name: "Joy Carr"}
	//employee.Street = "456 Elm St"
	employee.City = "Portland"
	//employee.State = "OR"
	//employee.PostalCode = "97222"
	//fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	//fmt.Println("State:", employee.State)
	//fmt.Println("Postal Code:", employee.PostalCode)
}
