package model

type Authentication struct {
	//Список разрешений, которые предоставляет данный токен к различным сервисам PayPal
	Scope string `json:"scope"`
	//Токен доступа
	AccessToken string `json:"access_token"`
	//Тип токена
	TokenType string `json:"token_type"`
	//Идентификатор приложения, которому выдан токен
	AppId string `json:"app_id"`
	//Время действия токена в секундах
	ExpiresIn int `json:"expires_in"`
	//Уникальный идентификатор запроса, который может быть использован для предотвращения повторного использования одного и того же запроса
	Nonce string `json:"nonce"`
}
