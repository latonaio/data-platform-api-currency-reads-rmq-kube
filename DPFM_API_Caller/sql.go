package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-currency-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-currency-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var currency *[]dpfm_api_output_formatter.Currency
	var currencyText *[]dpfm_api_output_formatter.CurrencyText
	for _, fn := range accepter {
		switch fn {
		case "Currency":
			func() {
				currency = c.Currency(mtx, input, output, errs, log)
			}()
		case "CurrencyText":
			func() {
				currencyText = c.CurrencyText(mtx, input, output, errs, log)
			}()
		case "CurrencyTexts":
			func() {
				currencyText = c.CurrencyTexts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Currency:     currency,
		CurrencyText: currencyText,
	}

	return data
}

func (c *DPFMAPICaller) Currency(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Currency {
	currency := input.Currency.Currency

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_currency_currency_data
		WHERE Currency = ?;`, currency,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCurrency(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) CurrencyText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.CurrencyText {
	var args []interface{}
	currency := input.Currency.Currency
	currencyText := input.Currency.CurrencyText

	cnt := 0
	for _, v := range currencyText {
		args = append(args, currency, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_currency_currency_text_data
		WHERE (Currency, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCurrencyText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) CurrencyTexts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.CurrencyText {
	var args []interface{}
	currencyText := input.Currency.CurrencyText

	cnt := 0
	for _, v := range currencyText {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_currency_currency_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToCurrencyTexts(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
