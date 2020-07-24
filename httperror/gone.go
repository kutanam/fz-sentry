package httperror

import "net/http"

// Gone is a constructor to create GoneError instance
func Gone(err error) error {
	return New(http.StatusGone, err)
}

// IsGoneError check whether given error is a GoneError
func IsGoneError(err error) bool {
	return GetInstance(err).Code == http.StatusGone
}
