package main

import (
	"Go_Head_first/go/src/github.com/headfirstgo/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	//date.year = 2019
	//date.month = 14
	//date.day = 50
	//fmt.Println(date)
	//date = calendar.Date{year: 0, month: 0, day: -2}
	//fmt.Println(date)

	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	event := calendar.Event{}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
