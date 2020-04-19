package player

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("player %s won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed\n", name)

			close(court)
			return
		}

		fmt.Printf("player %s hit %d\n", name, ball)
		ball++

		court <- ball
	}
}

// Run performs the 2 plays gorountines logic
func Run() {

	court := make(chan int)

	wg.Add(2)

	go player("simon", court)
	go player("bob", court)

	// Start the set
	court <- 1

	wg.Wait()

}
