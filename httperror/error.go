package httperror

import (
	"fmt"
	"net/http"
)

// Base is a struct that contain basic requirements for http error struct
type Base struct {
	Code       int    `json:"-"`
	StatusCode string `json:"code"`
	Detail     error  `json:"error"`
}

// Error is a function to implement error interface
func (e *Base) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Code, http.StatusText(e.Code), e.Detail.Error())
}

// GetCode is a function to return http error code
func (e *Base) GetCode() int {
	return e.Code
}

// GetStatusCode is a function to return error message
func (e *Base) GetStatusCode() string {
	return e.StatusCode
}

// Base is a constructor for http error with custom message
func New(code int, message string, err error) error {
	return &Base{
		Code:       code,
		StatusCode: message,
		Detail:     err,
	}
}

// Code is a constructor for http error with default status text message
func Code(code int, err error) error {
	return &Base{
		Code:       code,
		StatusCode: http.StatusText(code),
		Detail:     err,
	}
}
