package httperror

import "net/http"

// InternalServer is a constructor to create InternalServerError instance
func InternalServer(err error) Interface {
	return New(http.StatusInternalServerError, err)
}

// IsInternalServerError check whether given error is a InternalServerError
func IsInternalServerError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusInternalServerError
}
