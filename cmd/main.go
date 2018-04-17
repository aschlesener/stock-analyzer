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
	busiestDays := flag.Bool("busiestDays", false, "Specifies whether to calculate busiest days")
	biggestLoser := flag.Bool("biggestLoser", false, "Specifies whether to calculate the biggest loser")
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
	} else if *busiestDays {
		// calculate days where security's volume is >10% greater than average
		busiestDaysMap := analyzer.CalcBusiestDays(tickerMap)
		outputResult(busiestDaysMap)
	} else if *biggestLoser {
		// calulate security with longest streak of losing days
		loser := analyzer.CalcBiggestLoser(tickerMap)
		outputResult(loser)
	} else {
		// calculate monthly open/close averages for each security
		monthlyAverages := analyzer.CalcAverageMonthly(tickerMap)
		outputResult(monthlyAverages)
	}
}

// helper function to pretty-print JSON result to the command line
func outputResult(input interface{}) {
	b, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.Write(b)
	writer.WriteString("\n")
}
