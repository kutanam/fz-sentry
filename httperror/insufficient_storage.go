package httperror

import "net/http"

// InsufficientStorage is a constructor to create InsufficientStorageError instance
func InsufficientStorage(err error) Interface {
	return New(http.StatusInsufficientStorage, err)
}

// IsInsufficientStorageError check whether given error is a InsufficientStorageError
func IsInsufficientStorageError(err error) bool {
	return GetInstance(err).Code == http.StatusInsufficientStorage
}
