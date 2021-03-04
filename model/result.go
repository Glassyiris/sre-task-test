package model

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"request message"`
	Data    interface{} `json:"data" `
}
