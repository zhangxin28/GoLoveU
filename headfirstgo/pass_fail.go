// pass_fail reports whether a grade is passing or failing

package main

import (
	"fmt"
	"goloveu/headfirstgo/keyboard"
	"log"
)

// RunPassFail reports whether a grade is passing or failing
func RunPassFail() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	status := "failing"
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of ", grade, "is ", status)
}
