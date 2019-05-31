package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Define Feed struct
type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	// Load data file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// Defer to close file
	defer file.Close()
	// Decode json file
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
