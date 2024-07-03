package model

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}
