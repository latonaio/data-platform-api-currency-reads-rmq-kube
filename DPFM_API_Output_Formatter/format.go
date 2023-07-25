package dpfm_api_output_formatter

import (
	"data-platform-api-currency-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToCurrency(rows *sql.Rows) (*[]Currency, error) {
	defer rows.Close()
	currency := make([]Currency, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Currency{}

		err := rows.Scan(
			&pm.Currency,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &currency, nil
		}

		data := pm
		currency = append(currency, Currency{
			Currency: 				data.Currency,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &currency, nil
}

func ConvertToCurrencyText(rows *sql.Rows) (*[]CurrencyText, error) {
	defer rows.Close()
	currencyText := make([]CurrencyText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.CurrencyText{}

		err := rows.Scan(
			&pm.Currency,
			&pm.Language,
			&pm.CurrencyName,
			&pm.CurrencyLongName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &currencyText, err
		}

		data := pm
		currencyText = append(currencyText, CurrencyText{
			Currency:         		data.Currency,
			Language:         		data.Language,
			CurrencyName:     		data.CurrencyName,
			CurrencyLongName: 		data.CurrencyLongName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &currencyText, nil
}
