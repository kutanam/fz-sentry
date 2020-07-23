package httperror

import "net/http"

// Unauthorized is a constructor to create UnauthorizedError instance
func Unauthorized(err error) error {
	return Code(http.StatusUnauthorized, err)
}

// IsUnauthorizedError check whether given error is a UnauthorizedError
func IsUnauthorizedError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusUnauthorized
}
