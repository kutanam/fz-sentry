package httperror

import "net/http"

// Conflict is a constructor to create ConflictError instance
func Conflict(err error) error {
	return Code(http.StatusConflict, err)
}

// IsConflictError check whether given error is a ConflictError
func IsConflictError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusConflict
}
