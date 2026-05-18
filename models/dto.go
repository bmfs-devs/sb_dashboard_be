package models

type ResponseGeneral struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
	Data         any    `json:"data"`
}
