package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aschlesener/stock-analyzer/analyzer"
)

func main() {
	// parse command line input
	apiKey := flag.String("apiKey", "apiKey", "Your Quandl API key")
	maxDailyProfit := flag.Bool("maxDailyProfit", false, "Specifies whether to calculate maximum daily profit")
	flag.Parse()

	// call API to fetch and parse data
	tickerMap, err := analyzer.GetParsedData(*apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *maxDailyProfit {
		// calculate maximum daily profit for each security
		maximumDailyProfits := analyzer.CalcMaxDailyProfit(tickerMap)
		outputResult(maximumDailyProfits)
	} else {
		// calculate monthly open/close averages for each security
		monthlyAverages := analyzer.CalcAverageMonthly(tickerMap)
		outputResult(monthlyAverages)
	}
}

func outputResult(input interface{}) {
	b, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.Write(b)
}
