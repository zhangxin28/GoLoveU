// Package datafile allows reading data samples from file.
package datafile

import (
	"bufio"
	"os"
)

// GetStrings reads a string from each line of a file.
func GetStrings() ([]string, error) {
	var lines []string
	file, err := os.Open("votes.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
