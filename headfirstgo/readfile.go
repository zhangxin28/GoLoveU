// Read data from a txt file
package main

import (
	"bufio"
	"fmt"
	"goloveu/headfirstgo/datafile"
	"log"
	"os"
)

// RunReadFile reads content from a txt file
func RunReadFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// It is also possible that the bufio.Scanner encountered an error
	// while scanning through the file.
	// If it did, calling the Err method on the scanner
	// will return that err, which we log before exiting.
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	numbers, err := datafile.GetFloats()
	if err != nil {
		log.Fatal(err)
	}

	sum, average := averageandsum(numbers...)
	fmt.Printf("Sum = %.3f， Average = %.3f\n",
		sum, average)
}

// `...` means this function is a varadic function
func averageandsum(numbers ...float64) (float64, float64) {
	if len(numbers) == 0 {
		return 0, 0
	}

	sum := float64(0)
	for _, number := range numbers {
		sum += number
	}

	return sum, sum / float64(len(numbers))
}
