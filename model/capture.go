package model

import "time"

type Capture struct {
	Id                        string                     `json:"id"`
	Amount                    *Amount                    `json:"amount"`
	Status                    string                     `json:"status"`
	Note                      string                     `json:"note"`
	NoteToPayer               string                     `json:"note_to_payer"`
	PaymentInstruction        *PaymentInstruction        `json:"payment_instruction"`
	Payee                     *Payee                     `json:"payee"`
	SellerProtection          *SellerProtection          `json:"seller_protection"`
	FinalCapture              bool                       `json:"final_capture"`
	SellerReceivableBreakdown *SellerReceivableBreakdown `json:"seller_receivable_breakdown"`
	InvoiceId                 string                     `json:"invoice_id"`
	CreateTime                time.Time                  `json:"create_time"`
	UpdateTime                time.Time                  `json:"update_time"`
	Links                     []*Link                    `json:"links"`
}
