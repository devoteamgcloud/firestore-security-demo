package models

import (
	"fmt"
)

const (
	CodeOK                  Code = "OK"
	CodeCreated             Code = "Created"
	CodeUpdate              Code = "Updated"
	CodeMoved               Code = "Moved"
	CodeNotFound            Code = "NotFound"
	CodeUnprocessableEntity Code = "UnprocessableEntity"
	CodeServerSideError     Code = "ServerError"
)

type Code string

func (code Code) String() string {
	return string(code)
}

func (code Code) MarshalJSON() (text []byte, err error) {
	return []byte(fmt.Sprintf("\"%s\"", code)), nil
}

type Response[T any] struct {
	Code    Code `json:"code"              enums:"OK,Created,Updated,Moved,NotFound,UnprocessableEntity,ServerError"`
	Content T    `json:"content,omitempty"`
}

func (response *Response[T]) String() string {
	return fmt.Sprintf("%s response", response.Code)
}

type ErrorResponse struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}
