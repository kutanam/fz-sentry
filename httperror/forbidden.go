package httperror

import "net/http"

// Forbidden is a constructor to create ForbiddenError instance
func Forbidden(err error) error {
	return New(http.StatusForbidden, err)
}

// IsForbiddenError check whether given error is a ForbiddenError
func IsForbiddenError(err error) bool {
	return GetInstance(err).Code == http.StatusForbidden
}
