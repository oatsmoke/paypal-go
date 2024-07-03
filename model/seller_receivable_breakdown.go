package model

type SellerReceivableBreakdown struct {
	GrossAmount         *Amount        `json:"gross_amount"`
	PaypalFee           *Amount        `json:"paypal_fee"`
	NetAmount           *Amount        `json:"net_amount"`
	TotalRefundedAmount *Amount        `json:"total_refunded_amount"`
	PlatformFees        []*PlatformFee `json:"platform_fees"`
}
