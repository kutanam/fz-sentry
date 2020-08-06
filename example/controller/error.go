package controller

import (
	"errors"
	"net/http"

	"github.com/payfazz/fz-sentry/httperror"
	"github.com/payfazz/fz-sentry/logger"
	"github.com/payfazz/fz-sentry/loghttp"
	"go.uber.org/zap"
)

func Error() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.GetLogger(ctx)

		log.Debug("this is debug message")

		err := httperror.InternalServer(errors.New("undefined error"))
		log.Error(
			"this is error",
			zap.String("cause", err.Error()),
		)

		loghttp.Error(w, err)
	}
}
