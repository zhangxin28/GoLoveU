// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	floatBiteSize            = 64
	readStringDelim          = '\n'
	defaultFloatZero float64 = 0
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(readStringDelim)
	if err != nil {
		return defaultFloatZero, err
	}

	input = strings.TrimSpace(input)
	// here err is an assignment
	number, err := strconv.ParseFloat(input, floatBiteSize)
	if err != nil {
		return defaultFloatZero, err
	}

	return number, nil
}
