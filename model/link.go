package model

type Link struct {
	Rel     string `json:"rel"`
	Href    string `json:"href"`
	Method  string `json:"method"`
	EncType string `json:"encType"`
}
