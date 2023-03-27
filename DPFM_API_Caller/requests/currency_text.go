package requests

type CurrencyText struct {
	Currency         string  `json:"Currency"`
	Language         string  `json:"Language"`
	CurrencyName     *string `json:"CurrencyName"`
	CurrencyLongName *string `json:"CurrencyLongName"`
}
