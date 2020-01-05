package matchers

import (
	"goloveu/goinactioncode/matcherssearch/data"
)

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *data.Feed, searchTerm string) ([]*data.Result, error) {
	return nil, nil
}
