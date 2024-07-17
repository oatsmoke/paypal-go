package model

type Authentication struct {
	//List of permissions that this token grants to various PayPal services
	Scope string `json:"scope,omitempty"`
	//Access token
	AccessToken string `json:"access_token,omitempty"`
	//Token type
	TokenType string `json:"token_type,omitempty"`
	//Identifier of the application to which the token has been issued
	AppId string `json:"app_id,omitempty"`
	//Token validity time in seconds
	ExpiresIn int `json:"expires_in,omitempty"`
	//A unique request identifier that can be used to prevent reuse of the same request
	Nonce string `json:"nonce,omitempty"`
}
