package httperror

import "net/http"

// BadGateway is a constructor to create BadGatewayError instance
func BadGateway(err error) error {
	return Code(http.StatusBadGateway, err)
}

// IsBadGatewayError check whether given error is a BadGatewayError
func IsBadGatewayError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusBadGateway
}
