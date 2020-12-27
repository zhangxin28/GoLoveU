// tocelsius converts a temperature from Fahrenheit to Celsuis
package main

import (
	"fmt"
	"goloveu/headfirstgo/keyboard"
	"log"
)

// RunTocelsius converts a temperature from Fahrenheit to Celsuis
func RunTocelsius() {
	fmt.Print("Enter a temperature in Fahrenheit:")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsuis := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degress Celsuis.\n", celsuis)
}
