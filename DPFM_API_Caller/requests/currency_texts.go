package requests

type CurrencyTexts struct {
	Currency         string  `json:"Currency"`
	Language         string  `json:"Language"`
	CurrencyName     *string `json:"CurrencyName"`
	CurrencyLongName *string `json:"CurrencyLongName"`
}
