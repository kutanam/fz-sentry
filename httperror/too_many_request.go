package httperror

import "net/http"

// TooManyRequests is a constructor to create TooManyRequestsError instance
func TooManyRequests(err error) Interface {
	return New(http.StatusTooManyRequests, err)
}

// IsTooManyRequestsError check whether given error is a TooManyRequestsError
func IsTooManyRequestsError(err error) bool {
	return err != nil && GetInstance(err).GetCode() == http.StatusTooManyRequests
}
