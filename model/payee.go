package model

type Payee struct {
	MerchantId   string `json:"merchant_id"`
	EmailAddress string `json:"email_address"`
}
