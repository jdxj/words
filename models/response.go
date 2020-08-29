package models

import "time"

func NewResponse(code int, message string, data interface{}) *Response {
	resp := &Response{
		Code:      code,
		Timestamp: time.Now().Unix(),
		Message:   message,
		Data:      data,
	}
	return resp
}

// todo: 状态码定义
type Response struct {
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
