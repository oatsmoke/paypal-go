package model

type Customer struct {
	Id                 string `json:"id"`
	MerchantCustomerId string `json:"merchant_customer_id"`
}
