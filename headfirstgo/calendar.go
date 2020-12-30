package main

import (
	"fmt"
	"goloveu/headfirstgo/calendar"
	"log"
)

// RunCalendar tests the Encapsulation for a type
func RunCalendar() {
	event := calendar.Event{}

	err := event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(30)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetTitle("You win my boy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	fmt.Println(event.Date)
	fmt.Println(event.Title())
	fmt.Println(event)
}
