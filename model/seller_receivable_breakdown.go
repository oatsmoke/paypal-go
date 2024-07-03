package model

type SellerReceivableBreakdown struct {
	GrossAmount         *Amount        `json:"gross_amount,omitempty"`
	PaypalFee           *Amount        `json:"paypal_fee,omitempty"`
	NetAmount           *Amount        `json:"net_amount,omitempty"`
	TotalRefundedAmount *Amount        `json:"total_refunded_amount,omitempty"`
	PlatformFees        []*PlatformFee `json:"platform_fees,omitempty"`
}
