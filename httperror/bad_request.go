package httperror

import "net/http"

// BadRequest is a constructor to create BadRequestError instance
func BadRequest(err error) Interface {
	return New(http.StatusBadRequest, err)
}

// IsBadRequestError check whether given error is a BadRequestError
func IsBadRequestError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusBadRequest
}
