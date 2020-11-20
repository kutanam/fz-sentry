package httperror

import "net/http"

// Unauthorized is a constructor to create UnauthorizedError instance
func Unauthorized(err error) Interface {
	return New(http.StatusUnauthorized, err)
}

// IsUnauthorizedError check whether given error is a UnauthorizedError
func IsUnauthorizedError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusUnauthorized
}
