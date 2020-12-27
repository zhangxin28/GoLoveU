// guess chanllenges players to guess a random number.
package main

import (
	"fmt"
	"goloveu/headfirstgo/keyboard"
	"log"
	"math/rand"
	"time"

	anotherkeyboard "github.com/headfirstgo/keyboard"
)

// RunGuessGame is a game for user to guess a random number
func RunGuessGame() {
	// generate a integer from 0 to 100
	target := generateRandomIntNumber(100)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)
	guessresult := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		result, err := checkGuess(target)
		guessresult = result
		if err != nil {
			fmt.Println("Something wrong with it:")
			log.Println(err)
			continue
		}

		if result {
			fmt.Printf("Well done. %d times You got the right answer.\n",
				guesses)
			break
		}
	}
	if !guessresult {
		fmt.Printf("Sorry. You didn't guess the right anser: %d.\n",
			target)
	}
}

func guess() (float64, error) {
	return anotherkeyboard.GetFloat()
}

func checkGuess(target int) (result bool, err error) {
	result = false
	guess, err := keyboard.GetFloat()
	if err != nil {
		return
	}
	guessint := int(guess)
	if guessint < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else if guessint > target {
		fmt.Println("Oops. Your guess was HIGH.")
	} else {
		result = true
	}
	return
}

// It panics if n <= 0.
func generateRandomIntNumber(n int) int {
	if n <= 0 {
		panic("n must great than zero")
	}

	// get the current date and time, as as integer
	// call the Unix method on Time, which will convert it
	// to an integer. Specifically, it will convert it to
	// Unit time format, which is an integer with the number
	// of seconds since January 1, 1970.
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	return rand.Intn(n) + 1
}
