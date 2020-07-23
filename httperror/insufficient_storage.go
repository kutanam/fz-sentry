package httperror

import "net/http"

// InsufficientStorage is a constructor to create InsufficientStorageError instance
func InsufficientStorage(err error) error {
	return Code(http.StatusInsufficientStorage, err)
}

// IsInsufficientStorageError check whether given error is a InsufficientStorageError
func IsInsufficientStorageError(err error) bool {
	base, ok := err.(*Base)
	return ok && base.Code == http.StatusInsufficientStorage
}
