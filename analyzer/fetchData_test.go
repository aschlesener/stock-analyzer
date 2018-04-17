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
