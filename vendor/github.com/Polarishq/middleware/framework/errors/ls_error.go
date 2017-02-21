package errors

import (
	"encoding/json"
)

// IError general error interface
type IError interface {
	Code() int32
	Error() string
}

// ErrorAsJSON parses IError into jason
func ErrorAsJSON(err IError) []byte {
	b, _ := json.Marshal(struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}{err.Code(), err.Error()})
	return b
}

// LSError http error struct
type LSError struct {
	code int32
	msg  string
}

// Code returns http code
func (p LSError) Code() int32 {
	return p.code
}

// Error returns message associated with http code
func (p LSError) Error() string {
	return p.msg
}

// NewLSError creates new log service error
func NewLSError(code int32, msg string) IError {
	return &LSError{code, msg}
}