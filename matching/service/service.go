package service

import "time"

type Request struct {
	Method  string
	Headers map[string][]string
	Body    *DataBody
}

type Response struct {
	Code        int
	CodeMessage string
	Error       error
	Body        *DataBody
}

type DataBody struct {
	Timestamp time.Time
	ID        *int
	Data      *string
}

type Service interface {
	MakeRequest(*Request) *Response
}
