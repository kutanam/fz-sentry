package httperror

import "net/http"

// Conflict is a constructor to create ConflictError instance
func Conflict(err error) Interface {
	return New(http.StatusConflict, err)
}

// IsConflictError check whether given error is a ConflictError
func IsConflictError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusConflict
}
