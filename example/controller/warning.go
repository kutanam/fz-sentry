package controller

import (
	"net/http"

	"github.com/payfazz/fz-sentry/loghttp"

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

		loghttp.Write(w, `{"success": true}`, http.StatusOK)
	}
}
