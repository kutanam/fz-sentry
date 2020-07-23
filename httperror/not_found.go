package httperror

import "net/http"

// NotFound is a constructor to create NotFoundError instance
func NotFound(err error) error {
	return Code(http.StatusNotFound, err)
}

// IsNotFoundError check whether given error is a NotFoundError
func IsNotFoundError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusNotFound
}
