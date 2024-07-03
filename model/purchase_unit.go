package model

type PurchaseUnit struct {
	ReferenceId string    `json:"reference_id,omitempty"`
	Amount      *Amount   `json:"amount,omitempty"`
	Payee       *Payee    `json:"payee,omitempty"`
	Payments    *Payments `json:"payments,omitempty"`
}
