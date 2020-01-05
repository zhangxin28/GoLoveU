package matcherssearch

import (
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// Run the matchers search
func Run() {
	// Perform the search for the specified term.
	//search.Run("president")
}
