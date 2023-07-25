package requests

type CurrencyText struct {
	Currency         	string  `json:"Currency"`
	Language        	string  `json:"Language"`
	CurrencyName     	string  `json:"CurrencyName"`
	CurrencyLongName	*string `json:"CurrencyLongName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
