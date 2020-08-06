package httperror

import (
	"fmt"
	"net/http"
)

type Interface interface {
	Error() string
	RawError() string
	CompleteError() string
	GetDetail() error
	GetMessage() string
	GetCode() int
	SetMessage(message string)
}

// Base is a struct that contain basic requirements for http error struct
type Base struct {
	Code       int    `json:"-"`
	StatusCode string `json:"code"`
	Detail     error  `json:"-"`
	Message    string `json:"message"`
}

// Error implement error interface, and return error Message
func (e *Base) Error() string {
	if "" == e.Message {
		return e.RawError()
	}
	return e.Message
}

// RawError return basic error message
func (e *Base) RawError() string {
	return e.Detail.Error()
}

// CompleteError return complete error message with http code and base error message
func (e *Base) CompleteError() string {
	return fmt.Sprintf("%d %s: %v", e.Code, http.StatusText(e.Code), e.Detail)
}

// GetDetail get error detail from httperror.Base instance
func (e *Base) GetDetail() error {
	return e.Detail
}

// GetMessage get message from httperror.Base instance
func (e *Base) GetMessage() string {
	return e.Message
}

// GetCode get http status code from httperror.Base instance
func (e *Base) GetCode() int {
	return e.Code
}

// SetMessage set error message returned by this error instance
func (e *Base) SetMessage(message string) {
	e.Message = message
}

// Base constructor for http error with custom message
func New(code int, err error) Interface {
	if nil == err {
		return nil
	}

	if base := GetInstance(err); nil != base {
		err = base.GetDetail()
	}

	return &Base{
		Code:       code,
		StatusCode: http.StatusText(code),
		Detail:     err,
		Message:    err.Error(),
	}
}

// GetInstance get Base error instance from error interface, will return wrapped error with 500 http code on non Base error
func GetInstance(err error) Interface {
	if result, ok := err.(*Base); ok {
		return result
	}
	return &Base{
		Code:       500,
		StatusCode: http.StatusText(500),
		Detail:     err,
		Message:    err.Error(),
	}
}
