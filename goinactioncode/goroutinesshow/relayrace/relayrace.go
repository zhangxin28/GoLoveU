package relayrace

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func raceRun(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("runner %d running with baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("runner %d to the line\n", newRunner)
		go raceRun(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("runner %d finished, race over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("runner %d exchange with runner %d\n", runner, newRunner)

	baton <- newRunner
}

// Run performs to simulate a person running in the relay race
func Run() {
	baton := make(chan int)

	wg.Add(1)

	go raceRun(baton)

	baton <- 1

	wg.Wait()
}
