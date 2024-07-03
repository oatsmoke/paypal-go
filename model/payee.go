package model

type Payee struct {
	MerchantId   string `json:"merchant_id,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}
