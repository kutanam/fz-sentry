package logmiddleware

import (
	"context"
	"net/http"

	"github.com/payfazz/fz-sentry/logger"
	"github.com/payfazz/fz-sentry/loghttp"
	"go.uber.org/zap"
)

func DoHTTP(
	next http.HandlerFunc,
	before func(ctx context.Context, log *zap.Logger, r *http.Request) error,
	after func(ctx context.Context, log *zap.Logger, out []byte) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx := r.Context()
		log := logger.GetLogger(ctx)

		err = before(ctx, log, r)
		if nil != err {
			loghttp.Error(w, err)
			return
		}

		wr := loghttp.WrapWriter(w)

		next(wr, r)

		err = after(ctx, log, wr.Body)
		if nil != err {
			loghttp.Error(w, err)
			return
		}
	}
}
