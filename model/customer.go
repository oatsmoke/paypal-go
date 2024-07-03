package model

type Customer struct {
	Id                 string `json:"id,omitempty"`
	MerchantCustomerId string `json:"merchant_customer_id,omitempty"`
}
