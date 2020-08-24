package logger

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func GetZapPayloadField(payload interface{}) zap.Field {
	by, _ := json.Marshal(payload)
	return zap.String("payload", string(by))
}
