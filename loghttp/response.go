package loghttp

import (
	"encoding/json"
	"net/http"

	"github.com/payfazz/fz-sentry/httperror"
)

func Error(w http.ResponseWriter, err error) {
	be := httperror.GetInstance(err)
	Write(w, be, be.GetCode())
}

func Write(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
