package model

type PaymentInstruction struct {
	PlatformFees []*PlatformFee `json:"platform_fees"`
}
