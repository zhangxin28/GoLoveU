package main

import (
	"fmt"
)

func main() {
	fmt.Println("Head First Go!")
	// RunPassFail()
	// RunGuessGame()
	// RunTocelsius()
	// RunReadFile()
	// RunCount()
	// RunCountWithMap()
	// RunMagazine()
	RunReceiver()

}

func interestFunc() {
	var x int
	inc := func() int {
		x++
		return x
	}
	fmt.Println(func() (a, b int) {
		return inc(), inc()
	}())
}
