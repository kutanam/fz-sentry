package httperror

import "net/http"

// ServiceUnavailable is a constructor to create ServiceUnavailableError instance
func ServiceUnavailable(err error) Interface {
	return New(http.StatusServiceUnavailable, err)
}

// IsServiceUnavailableError check whether given error is a ServiceUnavailableError
func IsServiceUnavailableError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusServiceUnavailable
}
