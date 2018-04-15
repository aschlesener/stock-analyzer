package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aschlesener/stock-analyzer/analyzer"
)

func main() {
	// parse command line input
	apiKey := flag.String("apiKey", "apiKey", "Your Quandl API key")
	flag.Parse()

	// call API to fetch and parse data
	tickerMap, err := analyzer.GetParsedData(*apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	// calculate monthly open/close averages for each security
	monthlyAverages := analyzer.CalcAverageMonthly(tickerMap)

	// output JSON result for monthly open/close averages
	b, err := json.MarshalIndent(monthlyAverages, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Stdout.Write(b)
}
