package httperror

import "net/http"

// Forbidden is a constructor to create ForbiddenError instance
func Forbidden(err error) error {
	return Code(http.StatusForbidden, err)
}

// IsForbiddenError check whether given error is a ForbiddenError
func IsForbiddenError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusForbidden
}
