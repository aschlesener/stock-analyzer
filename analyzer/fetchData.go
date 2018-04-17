package analyzer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	apiEndpoint      = "https://www.quandl.com/api/v3/datatables/WIKI/PRICES"
	defaultStartDate = "2017-01-01"
	defaultEndDate   = "2017-06-30"
)

var (
	defaultTickers = []string{"COF", "GOOGL", "MSFT"}
)

// GetParsedData calls stock API to fetch data and returns parsed version of that data
func GetParsedData(apiKey string) (map[string][]DailyStockData, error) {
	startDate := defaultStartDate
	endDate := defaultEndDate
	tickers := defaultTickers

	// build API URL
	urlString := buildAPIURL(tickers, startDate, endDate, apiKey)

	// call API
	var client = &http.Client{
		Timeout: time.Second * 60,
	}

	response, err := client.Get(urlString)
	if err != nil {
		return nil, errors.New("Error fetching from API")
	}

	if response.StatusCode != 200 {
		if response.StatusCode == 400 {
			return nil, errors.New("Unauthorized - check your API key")
		}
		return nil, errors.New("Non-200 response: " + strconv.Itoa(response.StatusCode))
	}

	// parse API response
	defer response.Body.Close()

	var jsonResponse ApiResponse
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error parsing API response")
	}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, errors.New("Error parsing API response")
	}

	// convert JSON response to map of tickers with array of daily stock data
	tickerMap := convertResponse(jsonResponse)
	return tickerMap, nil
}

func buildAPIURL(tickers []string, startDate string, endDate string, apiKey string) string {
	url, _ := url.Parse(apiEndpoint)
	query := url.Query()
	tickersStr := strings.Join(tickers, ",")
	query.Set("date.gte", startDate)
	query.Set("date.lte", endDate)
	query.Set("ticker", tickersStr)
	query.Set("api_key", apiKey)
	url.RawQuery = query.Encode()
	urlString := url.String()

	return urlString
}

// handle converting array of interfaces to struct. The Quandl API returns data in a mixed form without json tags
func convertResponse(jsonResponse ApiResponse) map[string][]DailyStockData {
	columnNames := jsonResponse.Datatable.ColumnNames
	tickerMap := make(map[string][]DailyStockData)

	for _, data := range jsonResponse.Datatable.Results {
		s := DailyStockData{}
		for i, value := range data {
			// match column with column name from column names given in API response
			if i >= len(columnNames) {
				continue
			}
			switch strings.ToLower(columnNames[i].Name) {
			case "ticker":
				s.Name = value.(string)
			case "date":
				s.Date = value.(string)
			case "open":
				s.Open = value.(float64)
			case "high":
				s.High = value.(float64)
			case "low":
				s.Low = value.(float64)
			case "close":
				s.Close = value.(float64)
			case "volume":
				s.Volume = value.(float64)
			}
		}

		// place data in map to separate ticker data
		stockDatas, ok := tickerMap[s.Name]
		if !ok {
			tickerMap[s.Name] = []DailyStockData{s}
		} else {
			tickerMap[s.Name] = append(stockDatas, s)
		}
	}

	return tickerMap
}
