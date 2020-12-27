// Package datafile allows reading data samples from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

const floatBiteSize = 64

// GetFloats returns float array from a file
func GetFloats() ([]float64, error) {
	var numbers []float64
	file, err := os.Open("floatsData.txt")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(),
			floatBiteSize)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
