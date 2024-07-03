package model

type ExperienceContext struct {
	BrandName string `json:"brand_name,omitempty"`
	Locale    string `json:"locale,omitempty"`
	ReturnUrl string `json:"return_url,omitempty"`
	CancelUrl string `json:"cancel_url,omitempty"`
}
