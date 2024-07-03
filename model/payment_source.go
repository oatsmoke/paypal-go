package model

type PaymentSource struct {
	Card  *Card  `json:"card"`
	Token *Token `json:"token"`
}
