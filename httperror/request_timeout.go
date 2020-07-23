package httperror

import "net/http"

// RequestTimeout is a constructor to create RequestTimeoutError instance
func RequestTimeout(err error) error {
	return Code(http.StatusRequestTimeout, err)
}

// IsRequestTimeoutError check whether given error is a RequestTimeoutError
func IsRequestTimeoutError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusRequestTimeout
}
