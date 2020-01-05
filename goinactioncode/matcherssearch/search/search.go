package search

import (
	"goloveu/goinactioncode/matcherssearch/data"
	"goloveu/goinactioncode/matcherssearch/matchers"
	"log"
	"sync"
)

// Run performs the search logic.
func Run(searchTerm string) {
	// Retrive the list of feeds to search through.
	feeds, err := data.RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *data.Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		matcher := matchers.Retrieve(feed)

		// Launch the goroutine to perform the search.
		go func(matcher matchers.Matcher, feed *data.Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to siginal to the Display function
		// that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are availabel and return
	// after the final result is displayed.
	Display(results)
}
