package analyzer

type ApiResponse struct {
	Datatable Data `json:"datatable"`
}

type Data struct {
	Results     [][]interface{} `json:"data"`
	ColumnNames []ColumnName    `json:"columns"`
}

type ColumnName struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DailyStockData struct {
	Name   string
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type AverageMonthlyPrices struct {
	Month        string  `json:"month"`
	AverageOpen  float64 `json:"average_open"`
	AverageClose float64 `json:"average_close"`
}
