package httperror

import "net/http"

// InternalServer is a constructor to create InternalServerError instance
func InternalServer(err error) error {
	return New(http.StatusInternalServerError, err)
}

// IsInternalServerError check whether given error is a InternalServerError
func IsInternalServerError(err error) bool {
	return GetInstance(err).Code == http.StatusInternalServerError
}
