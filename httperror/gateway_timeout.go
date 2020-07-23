package httperror

import "net/http"

// GatewayTimeout is a constructor to create GatewayTimeoutError instance
func GatewayTimeout(err error) error {
	return Code(http.StatusGatewayTimeout, err)
}

// IsGatewayTimeoutError check whether given error is a GatewayTimeoutError
func IsGatewayTimeoutError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusGatewayTimeout
}
