package httperror

import "net/http"

// UnprocessableEntity is a constructor to create UnprocessableEntityError instance
func UnprocessableEntity(err error) Interface {
	return New(http.StatusUnprocessableEntity, err)
}

// IsUnprocessableEntityError check whether given error is a UnprocessableEntityError
func IsUnprocessableEntityError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusUnprocessableEntity
}
