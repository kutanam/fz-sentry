package logger

import (
	"encoding/json"
	"net/http"
)

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func GetPayloadString(payload interface{}) string {
	by, _ := json.Marshal(payload)
	return string(by)
}
