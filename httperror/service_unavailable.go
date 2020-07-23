package httperror

import "net/http"

// ServiceUnavailable is a constructor to create ServiceUnavailableError instance
func ServiceUnavailable(err error) error {
	return Code(http.StatusServiceUnavailable, err)
}

// IsServiceUnavailableError check whether given error is a ServiceUnavailableError
func IsServiceUnavailableError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusServiceUnavailable
}
