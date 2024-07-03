package model

type PayPal struct {
	Id            string          `json:"id"`
	Customer      *Customer       `json:"customer"`
	Status        string          `json:"status"`
	Intent        string          `json:"intent"`
	PaymentSource *PaymentSource  `json:"payment_source"`
	PurchaseUnits []*PurchaseUnit `json:"purchase_units"`
	Links         []*Link         `json:"links"`
}
