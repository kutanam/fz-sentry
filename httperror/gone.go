package httperror

import "net/http"

// Gone is a constructor to create GoneError instance
func Gone(err error) error {
	return Code(http.StatusGone, err)
}

// IsGoneError check whether given error is a GoneError
func IsGoneError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusGone
}
