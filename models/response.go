package models

// todo: 状态码定义
type Response struct {
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
