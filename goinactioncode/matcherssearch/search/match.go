package search

import (
	"goloveu/goinactioncode/matcherssearch/data"
	"goloveu/goinactioncode/matcherssearch/matchers"
	"log"
)

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher matchers.Matcher, feed *data.Feed, searchTerm string, results chan<- *data.Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the individual goroutines.
func Display(results chan *data.Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
