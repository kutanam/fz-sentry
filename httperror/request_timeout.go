package httperror

import "net/http"

// RequestTimeout is a constructor to create RequestTimeoutError instance
func RequestTimeout(err error) Interface {
	return New(http.StatusRequestTimeout, err)
}

// IsRequestTimeoutError check whether given error is a RequestTimeoutError
func IsRequestTimeoutError(err error) bool {
	return GetInstance(err).Code == http.StatusRequestTimeout
}
