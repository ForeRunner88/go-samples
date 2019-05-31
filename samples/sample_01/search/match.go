package search

import (
	"log"
)

// Define Result struct
type Result struct {
	Field   string
	Content string
}

// Matcher定义所有搜索类型共同的行为
// Use Interface
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Use goroutine to run Match function to get search result
// Put result into chan
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// use matcher to search
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	// put result in chan
	for _, result := range searchResults {
		results <- result
	}
}

// Display function will get the search result from each goroutine
// Print the result onto console
func Display(results chan *Result) {
	// The chan will be blocked until some result is putted in
	// If the chan is closed, the for cycle will break
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
