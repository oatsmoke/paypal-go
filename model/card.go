package model

type Card struct {
	Number            string             `json:"number,omitempty"`
	Expiry            string             `json:"expiry,omitempty"`
	Name              string             `json:"name,omitempty"`
	VaultId           string             `json:"vault_id,omitempty"`
	BillingAddress    *BillingAddress    `json:"billing_address,omitempty"`
	ExperienceContext *ExperienceContext `json:"experience_context,omitempty"`
	Attributes        *Attributes        `json:"attributes,omitempty"`
	Brand             string             `json:"brand,omitempty"`
	LastDigits        string             `json:"last_digits,omitempty"`
}
