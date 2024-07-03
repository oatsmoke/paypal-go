package model

type PurchaseUnit struct {
	ReferenceId string    `json:"reference_id"`
	Amount      *Amount   `json:"amount"`
	Payee       *Payee    `json:"payee"`
	Payments    *Payments `json:"payments"`
}
