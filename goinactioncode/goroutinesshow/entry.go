package gorountinesshow

import (
	"goloveu/goinactioncode/goroutinesshow/chars"
	"goloveu/goinactioncode/goroutinesshow/player"
	"goloveu/goinactioncode/goroutinesshow/prime"
	"goloveu/goinactioncode/goroutinesshow/relayrace"
	"log"
	"os"
)

// GoRountinesShow preforms the sample for matchers search
type GoRountinesShow struct{}

// Run implements the behavior for the goinactioncode Run.
func (gr GoRountinesShow) Run() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)

	// Perform the player to Run
	player.Run()

	// Perform the chars to Run
	chars.Run()

	// Perform the prime to Run
	prime.Run()

	// Perform the ralayrace to Run
	relayrace.Run()
}
