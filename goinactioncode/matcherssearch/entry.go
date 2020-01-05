package matcherssearch

import (
	"goloveu/goinactioncode"
	"goloveu/goinactioncode/matcherssearch/search"
	"log"
	"os"
)

// MatchersSearch preforms the sample for matchers search
type MatchersSearch struct{}

// Run implements the behavior for the goinactioncode Run.
func (m MatchersSearch) Run() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)

	// Perform the search for the specified term.
	search.Run("president")
}

// init perfomrs to init the
func init() {
	var executer MatchersSearch
	goinactioncode.RegisterExecuter("MatchersSearch", executer)
}
