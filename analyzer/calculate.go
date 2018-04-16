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

// CalcMaxDailyProfit calculates the maximum daily profit for each security
func CalcMaxDailyProfit(tickerMap map[string][]DailyStockData) []MaxDailyProfit {
	dailyProfits := make([]MaxDailyProfit, 0)
	for ticker, dailyDatas := range tickerMap {
		maxDailyProfit := MaxDailyProfit{Ticker: ticker}

		// find date that provides the maximum daily profit for buying high and selling low
		maxProfit := 0.0
		var maxProfitDate string
		for _, dailyData := range dailyDatas {
			dailyProfit := dailyData.High - dailyData.Low
			if dailyProfit > maxProfit {
				maxProfit = dailyProfit
				maxProfitDate = dailyData.Date
			}
		}
		maxDailyProfit.Profit = maxProfit
		maxDailyProfit.Date = maxProfitDate
		dailyProfits = append(dailyProfits, maxDailyProfit)
	}

	return dailyProfits
}

// CalcBusiestDays calculates the days where a security's volume was over 10% greater than the average volume for that security
func CalcBusiestDays(tickerMap map[string][]DailyStockData) map[string]BusiestDays {
	busiestDaysMap := make(map[string]BusiestDays)

	for ticker, dailyDatas := range tickerMap {
		var busiestDays BusiestDays

		// calculate average volume for this security
		volumeSum := 0.0
		for _, dailyData := range dailyDatas {
			volumeSum += dailyData.Volume
		}
		busiestDays.AverageVolume = volumeSum / float64(len(dailyDatas))

		// find days where volume is over 10% greater than the average volume
		days := make([]BusiestDay, 0)
		threshold := (.10 * busiestDays.AverageVolume) + busiestDays.AverageVolume
		for _, dailyData := range dailyDatas {
			if dailyData.Volume > threshold {
				day := BusiestDay{Date: dailyData.Date, Volume: dailyData.Volume}
				days = append(days, day)
			}
		}
		busiestDays.Days = days
		busiestDaysMap[ticker] = busiestDays
	}

	return busiestDaysMap
}

// CalcBiggestLoser caculates the security that had the most days where the closing price was less than the opening price
func CalcBiggestLoser(tickerMap map[string][]DailyStockData) BiggestLoser {
	losingDaysMap := make(map[string]int64)

	// calculate number of loss days for each security
	for ticker, dailyDatas := range tickerMap {
		var numLossDays int64
		for _, dailyData := range dailyDatas {
			if dailyData.Close < dailyData.Open {
				numLossDays++
			}
		}
		losingDaysMap[ticker] = numLossDays
	}

	// determine which security had greatest number of loss days
	var maxLossDays int64
	var maxLossTicker string
	for ticker, numLossDays := range losingDaysMap {
		if numLossDays > maxLossDays {
			maxLossDays = numLossDays
			maxLossTicker = ticker
		}
	}

	return BiggestLoser{Ticker: maxLossTicker, NumberDaysLoser: maxLossDays}
}
