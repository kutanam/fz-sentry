package controller

import (
	"net/http"

	"github.com/payfazz/fz-sentry/logger"
	"go.uber.org/zap"
)

func Warning() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.GetLogger(ctx)

		log.Debug("this is debug message")

		log.Warn(
			"this is warning",
			zap.String("cause", "warning occured"),
		)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)

		_, _ = w.Write([]byte("warning"))
	}
}
