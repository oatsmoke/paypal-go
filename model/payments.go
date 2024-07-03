package model

type Payments struct {
	Captures []*Capture `json:"captures,omitempty"`
}
