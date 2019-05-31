package search

// Define default matcher
type defaultMatcher struct{}

// Register defaultMatcher
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search defaultMatcher method
// Default search nothing
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
