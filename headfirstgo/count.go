// count tallies the number of times each line
// occurs within a file
package main

import (
	"fmt"
	"goloveu/headfirstgo/datafile"
	"log"
	"sort"
)

// RunCount tallies the number of times each line occurs within a file
func RunCount() {
	lines, err := datafile.GetStrings()
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}

		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

// RunCountWithMap tallies the number of times each line occurs within a file
func RunCountWithMap() {
	lines, err := datafile.GetStrings()
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	// The for...range loop processes map keys and values
	// in a random order because map is an unordered collection
	// of keys and values.
	// When you use a for...range loop with a map, you never know
	// what order you'll get the map's contents in.
	// If you need more consistent ordering,
	// you'll need to write the code for that yourself.
	for name, count := range counts {
		fmt.Printf("%s: %d\n", name, count)
	}

	// below code to do the ordering
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, counts[name])
	}
}
