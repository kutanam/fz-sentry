package httperror

import "net/http"

// TooManyRequests is a constructor to create TooManyRequestsError instance
func TooManyRequests(err error) error {
	return Code(http.StatusTooManyRequests, err)
}

// IsTooManyRequestsError check whether given error is a TooManyRequestsError
func IsTooManyRequestsError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusTooManyRequests
}
