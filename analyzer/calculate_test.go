package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcAverageMonthly(t *testing.T) {
	sampleTickerMap := createSampleTickerMap()
	averageMonthlyPrices := CalcAverageMonthly(sampleTickerMap)

	assert.Equal(t, 1, len(averageMonthlyPrices["COF"]))
	assert.Equal(t, "2017-01", averageMonthlyPrices["COF"][0].Month)
	assert.Equal(t, 88.7475, averageMonthlyPrices["COF"][0].AverageClose)
	assert.Equal(t, 88.935, averageMonthlyPrices["COF"][0].AverageOpen)

	assert.Equal(t, 1, len(averageMonthlyPrices["GOOGL"]))
	assert.Equal(t, "2017-01", averageMonthlyPrices["GOOGL"][0].Month)
	assert.Equal(t, 809.6, averageMonthlyPrices["GOOGL"][0].AverageClose)
	assert.Equal(t, 806.0033333333334, averageMonthlyPrices["GOOGL"][0].AverageOpen)
}

func TestCalcMaxDailyProfit(t *testing.T) {
	sampleTickerMap := createSampleTickerMap()
	maxDailyProfits := CalcMaxDailyProfit(sampleTickerMap)

	for _, maxDailyProfit := range maxDailyProfits {
		if maxDailyProfit.Ticker == "COF" {
			assert.Equal(t, "2017-01-05", maxDailyProfit.Date)
			assert.Equal(t, 2.329900000000009, maxDailyProfit.Profit)
		} else if maxDailyProfit.Ticker == "GOOGL" {
			assert.Equal(t, "2017-01-03", maxDailyProfit.Date)
			assert.Equal(t, 14.544999999999959, maxDailyProfit.Profit)
		}
	}
}

func TestCalcBusiestDays(t *testing.T) {
	sampleTickerMap := createSampleTickerMap()
	busiestDays := CalcBusiestDays(sampleTickerMap)

	assert.Equal(t, 2.60467275e+06, busiestDays["COF"].AverageVolume)
	assert.Equal(t, []BusiestDay{{"2017-01-03", 3.441067e+06}}, busiestDays["COF"].Days)

	assert.Equal(t, 1.604969e+06, busiestDays["GOOGL"].AverageVolume)
	assert.Equal(t, []BusiestDay{{"2017-01-03", 1.959033e+06}}, busiestDays["GOOGL"].Days)
}

func TestCalcBiggestLoser(t *testing.T) {
	sampleTickerMap := createSampleTickerMap()
	biggestLoser := CalcBiggestLoser(sampleTickerMap)

	assert.Equal(t, int64(2), biggestLoser.NumberDaysLoser)
	assert.Equal(t, "COF", biggestLoser.Ticker)
}

func createSampleTickerMap() map[string][]DailyStockData {
	sampleTickerMap := make(map[string][]DailyStockData)

	sampleTickerMap["COF"] = []DailyStockData{
		{"COF", "2017-01-03", 88.55, 89.6, 87.79, 88.87, 3.441067e+06},
		{"COF", "2017-01-04", 89.13, 90.77, 89.13, 90.3, 2.630905e+06},
		{"COF", "2017-01-05", 89.84, 89.9299, 87.6, 88.38, 2.223944e+06},
		{"COF", "2017-01-09", 88.22, 88.4, 87.4, 87.44, 2.122775e+06}}
	sampleTickerMap["GOOGL"] = []DailyStockData{
		{"GOOGL", "2017-01-03", 800.62, 811.435, 796.89, 808.01, 1.959033e+06},
		{"GOOGL", "2017-01-04", 809.89, 813.43, 804.11, 807.77, 1.515339e+06},
		{"GOOGL", "2017-01-05", 807.5, 813.74, 805.92, 813.02, 1.340535e+06}}

	return sampleTickerMap
}
