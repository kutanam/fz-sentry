package httperror

import "net/http"

// Unauthorized is a constructor to create UnauthorizedError instance
func Unauthorized(err error) error {
	return New(http.StatusUnauthorized, err)
}

// IsUnauthorizedError check whether given error is a UnauthorizedError
func IsUnauthorizedError(err error) bool {
	return GetInstance(err).Code == http.StatusUnauthorized
}
