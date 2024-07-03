package model

type Amount struct {
	CurrencyCode string `json:"currency_code,omitempty"`
	Value        string `json:"value,omitempty"`
}
