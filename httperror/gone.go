package httperror

import "net/http"

// Gone is a constructor to create GoneError instance
func Gone(err error) Interface {
	return New(http.StatusGone, err)
}

// IsGoneError check whether given error is a GoneError
func IsGoneError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusGone
}
