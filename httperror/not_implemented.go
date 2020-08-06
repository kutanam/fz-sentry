package httperror

import "net/http"

// NotImplemented is a constructor to create NotImplementedError instance
func NotImplemented(err error) Interface {
	return New(http.StatusNotImplemented, err)
}

// IsNotImplementedError check whether given error is a NotImplementedError
func IsNotImplementedError(err error) bool {
	return GetInstance(err).Code == http.StatusNotImplemented
}
