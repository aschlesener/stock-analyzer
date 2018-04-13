package main

import (
	"flag"
	"fmt"

	"github.com/aschlesener/stock-analyzer/analyzer"
)

func main() {
	// parse command line input
	apiKey := flag.String("apiKey", "apiKey", "Your Quandl API key")
	flag.Parse()

	// call API to fetch and parse data
	_, err := analyzer.GetParsedData(*apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
}
