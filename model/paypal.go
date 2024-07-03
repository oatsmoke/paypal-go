package model

type PayPal struct {
	Id            string          `json:"id,omitempty"`
	Customer      *Customer       `json:"customer,omitempty"`
	Status        string          `json:"status,omitempty"`
	Intent        string          `json:"intent,omitempty"`
	PaymentSource *PaymentSource  `json:"payment_source,omitempty"`
	PurchaseUnits []*PurchaseUnit `json:"purchase_units,omitempty"`
	Links         []*Link         `json:"links,omitempty"`
}
