package model

type PaymentSource struct {
	Card  *Card  `json:"card,omitempty"`
	Token *Token `json:"token,omitempty"`
}
