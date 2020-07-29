package controller

import (
	"errors"
	"net/http"

	"github.com/payfazz/fz-sentry/logger"
	"go.uber.org/zap"
)

func Error() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.GetLogger(ctx)

		log.Debug("this is debug message")

		err := errors.New("undefined error")
		log.Error(
			"this is error",
			zap.String("cause", err.Error()),
		)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)

		_, _ = w.Write([]byte("error"))
	}
}
