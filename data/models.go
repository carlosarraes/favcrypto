package data

type Currency struct {
	ID       int
	Favorite bool
	Name     string
	Symbol   string
	Price    float64
}

type Currencies []Currency

type CurrencyData struct {
	ID       int     `json:"id"`
	Favorite bool    `json:"favorite"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	Price    float64 `json:"price"`
}

type CurrenciesData []CurrencyData

func (c Currencies) ToCurrenciesData() CurrenciesData {
	currenciesData := CurrenciesData{}
	for _, currency := range c {
		currencyData := CurrencyData(currency)
		currenciesData = append(currenciesData, currencyData)
	}
	return currenciesData
}
