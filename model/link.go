package model

type Link struct {
	Rel     string `json:"rel,omitempty"`
	Href    string `json:"href,omitempty"`
	Method  string `json:"method,omitempty"`
	EncType string `json:"encType,omitempty"`
}
