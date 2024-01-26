package dto

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	IsError bool   `json:"isError"`
	Data    any    `json:"data,omitempty"`
}
