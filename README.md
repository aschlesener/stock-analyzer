
# stock-analyzer

StockAnalyzer calculates stock data using the [Quandl WIKI stock price API](https://www.quandl.com/databases/WIKIP/documentation/about).

Given a list of securities and a date range, StockAnalyzer will calculate the following information:
 - Average monthly open and close prices
 - Maximum daily profit
	 - Day within dataset that would produce the maximum daily profit, if purchased at that day's low and sold at that day's high
 - Busiest days
	- &gt;10% average volume for that security
 - Biggest loser
	 - Security with the most days where the closing price was less than the opening price

## Usage
Prerequisites: [Go](https://golang.org/dl/).
You also must have an API key from Quandl. You can get one by signing up for a free account at [Quandl](https://www.quandl.com)

```
go get github.com/aschlesener/stock-analzyer
cd $GOPATH/src/github.com/aschlesener/stock-analyzer/cmd
go build
./cmd -apiKey="YOUR QUANDL API KEY HERE"
```

## License
MIT
