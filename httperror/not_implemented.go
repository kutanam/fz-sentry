package httperror

import "net/http"

// NotImplemented is a constructor to create NotImplementedError instance
func NotImplemented(err error) error {
	return Code(http.StatusNotImplemented, err)
}

// IsNotImplementedError check whether given error is a NotImplementedError
func IsNotImplementedError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusNotImplemented
}
