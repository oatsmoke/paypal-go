package model

type Card struct {
	Number            string             `json:"number"`
	Expiry            string             `json:"expiry"`
	Name              string             `json:"name"`
	VaultId           string             `json:"vault_id"`
	BillingAddress    *BillingAddress    `json:"billing_address"`
	ExperienceContext *ExperienceContext `json:"experience_context"`
	Attributes        *Attributes        `json:"attributes"`
	Brand             string             `json:"brand"`
	LastDigits        string             `json:"last_digits"`
}
