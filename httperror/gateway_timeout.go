package httperror

import "net/http"

// GatewayTimeout is a constructor to create GatewayTimeoutError instance
func GatewayTimeout(err error) Interface {
	return New(http.StatusGatewayTimeout, err)
}

// IsGatewayTimeoutError check whether given error is a GatewayTimeoutError
func IsGatewayTimeoutError(err error) bool {
	return GetInstance(err).GetCode() == http.StatusGatewayTimeout
}
