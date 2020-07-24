package httperror

import (
	"fmt"
)

// Base is a struct that contain basic requirements for http error struct
type Base struct {
	Code    int    `json:"code"`
	Detail  error  `json:"error"`
	Message string `json:"message"`
}

// Error implement error interface
func (e *Base) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

// SetMessage set error message returned by this error instance
func (e *Base) SetMessage(message string) {
	e.Message = message
}

// Base constructor for http error with custom message
func New(code int, err error) error {
	if base := GetInstance(err); nil != base {
		err = base.Detail
	}

	return &Base{
		Code:   code,
		Detail: err,
	}
}

// GetInstance get Base error instance from error interface, will return wrapped error with 500 http code on non Base error
func GetInstance(err error) *Base {
	if result, ok := err.(*Base); ok {
		return result
	}
	return New(500, err).(*Base)
}
