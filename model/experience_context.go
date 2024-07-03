package model

type ExperienceContext struct {
	BrandName string `json:"brand_name"`
	Locale    string `json:"locale"`
	ReturnUrl string `json:"return_url"`
	CancelUrl string `json:"cancel_url"`
}
