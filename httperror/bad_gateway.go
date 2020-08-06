package httperror

import "net/http"

// BadGateway is a constructor to create BadGatewayError instance
func BadGateway(err error) Interface {
	return New(http.StatusBadGateway, err)
}

// IsBadGatewayError check whether given error is a BadGatewayError
func IsBadGatewayError(err error) bool {
	return GetInstance(err).GetCode() == http.StatusBadGateway
}
