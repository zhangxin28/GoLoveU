package gorountinesshow

import (
	"goloveu/goinactioncode/goroutinesshow/player"
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
}
