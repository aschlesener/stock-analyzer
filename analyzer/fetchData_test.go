package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildApiURL(t *testing.T) {
	tickers := []string{"ticker1", "ticker2", "ticker3"}
	apiKey := "myApiKey"
	apiURL := buildAPIURL(tickers, defaultStartDate, defaultEndDate, apiKey)

	assert.Equal(t, apiEndpoint+"?"+"api_key="+apiKey+"&"+"date.gte="+defaultStartDate+"&"+"date.lte="+
		defaultEndDate+"&"+"ticker="+"ticker1%2Cticker2%2Cticker3", apiURL)
}

func TestConvertResponse(t *testing.T) {
	sampleAPIResponse := createSampleAPIResponse()
	convertedResponse := convertResponse(sampleAPIResponse)

	assert.Equal(t, 1, len(convertedResponse))
	assert.Equal(t, 2, len(convertedResponse["COF"]))
	assert.Equal(t, "COF", convertedResponse["COF"][0].Name)
	assert.Equal(t, "2017-01-03", convertedResponse["COF"][0].Date)
	assert.Equal(t, 88.55, convertedResponse["COF"][0].Open)
	assert.Equal(t, 89.6, convertedResponse["COF"][0].High)
	assert.Equal(t, 87.79, convertedResponse["COF"][0].Low)
	assert.Equal(t, 88.87, convertedResponse["COF"][0].Close)
	assert.Equal(t, 3441067.0, convertedResponse["COF"][0].Volume)
}

func createSampleAPIResponse() ApiResponse {
	columnNames := []ColumnName{{Name: "ticker", Type: "String"},
		{Name: "date", Type: "Date"},
		{Name: "open", Type: "BigDecimal(34,12)"},
		{Name: "high", Type: "BigDecimal(34,12)"},
		{Name: "low", Type: "BigDecimal(34,12)"},
		{Name: "close", Type: "BigDecimal(34,12)"},
		{Name: "volume", Type: "BigDecimal(37,15)"},
		{Name: "ex-dividend", Type: "BigDecimal(42,20)"}}

	results := [][]interface{}{{"COF", "2017-01-03", 88.55, 89.6, 87.79, 88.87, 3441067.0, 0.0, 1.0, 87.30123847595, 88.336431027048, 86.551956248488, 87.616725729618, 3441067.0},
		{"COF", "2017-01-04", 89.13, 90.77, 89.13, 90.3, 2630905.0, 0.0, 1.0, 87.873059123223, 89.489931298271, 87.873059123223, 89.026559394447, 2630905.0}}

	data := Data{Results: results, ColumnNames: columnNames}
	apiResponse := ApiResponse{Datatable: data}

	return apiResponse
}
