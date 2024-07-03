package model

import "time"

type Capture struct {
	Id                        string                     `json:"id,omitempty"`
	Amount                    *Amount                    `json:"amount,omitempty"`
	Status                    string                     `json:"status,omitempty"`
	Note                      string                     `json:"note,omitempty"`
	NoteToPayer               string                     `json:"note_to_payer,omitempty"`
	PaymentInstruction        *PaymentInstruction        `json:"payment_instruction,omitempty"`
	Payee                     *Payee                     `json:"payee,omitempty"`
	SellerProtection          *SellerProtection          `json:"seller_protection,omitempty"`
	FinalCapture              bool                       `json:"final_capture,omitempty"`
	SellerReceivableBreakdown *SellerReceivableBreakdown `json:"seller_receivable_breakdown,omitempty"`
	InvoiceId                 string                     `json:"invoice_id,omitempty"`
	CreateTime                time.Time                  `json:"create_time,omitempty"`
	UpdateTime                time.Time                  `json:"update_time,omitempty"`
	Links                     []*Link                    `json:"links,omitempty"`
}
