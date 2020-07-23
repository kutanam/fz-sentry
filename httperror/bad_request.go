package httperror

import "net/http"

// BadRequest is a constructor to create BadRequestError instance
func BadRequest(err error) error {
	return Code(http.StatusBadRequest, err)
}

// IsBadRequestError check whether given error is a BadRequestError
func IsBadRequestError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusBadRequest
}
