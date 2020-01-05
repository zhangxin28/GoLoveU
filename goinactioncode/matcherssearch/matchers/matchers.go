package matchers

import (
	"goloveu/goinactioncode/matcherssearch/data"
	"log"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher)

// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *data.Feed, searchTerm string) ([]*data.Result, error)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

// Retrieve returns a matcher,if not match return the default matcher
func Retrieve(feed *data.Feed) Matcher {
	// Retrieve a matcher for the search.
	matcher, exists := matchers[feed.Type]
	if !exists {
		matcher = matchers["default"]
	}
	return matcher
}
