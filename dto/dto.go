package dto

// Response 通用回复结构体
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    uint        `json:"code"`
}
