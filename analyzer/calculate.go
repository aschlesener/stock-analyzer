package analyzer

// CalcAverageMonthly caclulates average monthly open and close prices for each security
func CalcAverageMonthly(tickerMap map[string][]DailyStockData) map[string][]AverageMonthlyPrices {
	tickerAveragesMap := make(map[string][]AverageMonthlyPrices)
	for ticker, dailyDatas := range tickerMap {
		averageMonthlyPrices := make([]AverageMonthlyPrices, 0)

		// group data for each ticker by month
		monthMap := make(map[string][]DailyStockData)
		for _, dailyData := range dailyDatas {
			// get month from date - 2017-01-03 becomes 2017-01
			month := dailyData.Date[:7]

			stockDatas, ok := monthMap[month]
			if !ok {
				monthMap[month] = []DailyStockData{dailyData}
			} else {
				monthMap[month] = append(stockDatas, dailyData)
			}
		}

		// calculate averages for each month
		for month, datas := range monthMap {
			monthlyPrices := AverageMonthlyPrices{Month: month}
			sumOpen := 0.0
			sumClose := 0.0
			for _, data := range datas {
				sumOpen += data.Open
				sumClose += data.Close
			}
			monthlyPrices.AverageClose = sumClose / float64(len(datas))
			monthlyPrices.AverageOpen = sumOpen / float64(len(datas))
			averageMonthlyPrices = append(averageMonthlyPrices, monthlyPrices)
		}

		tickerAveragesMap[ticker] = averageMonthlyPrices
	}
	return tickerAveragesMap
}
