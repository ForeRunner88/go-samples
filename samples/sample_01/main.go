package main

import (
	"log"
	"os"

	_ "github.com/ForeRunner88/go-samples/samples/sample_01/matchers"
	"github.com/ForeRunner88/go-samples/samples/sample_01/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
