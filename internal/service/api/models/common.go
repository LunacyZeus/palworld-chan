package models

type Response struct {
	Code    int    `json:"code"`
	Result  any    `json:"result"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
