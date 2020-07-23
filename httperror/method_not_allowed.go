package httperror

import "net/http"

// MethodNotAllowed is a constructor to create MethodNotAllowedError instance
func MethodNotAllowed(err error) error {
	return Code(http.StatusMethodNotAllowed, err)
}

// IsMethodNotAllowedError check whether given error is a MethodNotAllowedError
func IsMethodNotAllowedError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusMethodNotAllowed
}
