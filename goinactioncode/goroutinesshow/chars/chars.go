package chars

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func lettershow(lower bool) {
	defer wg.Done()

	char := 'a'
	if !lower {
		char = 'A'
	}

	for count := 0; count < 3; count++ {
		for charIn := char; charIn < char+26; charIn++ {
			fmt.Printf("%c", charIn)
		}
		fmt.Println("")
	}
}

// Run performs the 2 plays gorountines logic for lettershow
func Run() {
	runtime.GOMAXPROCS(2)

	wg.Add(2)

	fmt.Println("starting goroutines")

	go lettershow(true)
	go lettershow(false)

	fmt.Println("wait to finish")
	wg.Wait()

	fmt.Println("\nterminating program")
}
