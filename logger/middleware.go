package logger

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/payfazz/fz-sentry/loghttp"
	"go.uber.org/zap"
)

func DoGRPC(
	f endpoint.Endpoint,
	before func(ctx context.Context, log *zap.Logger, in interface{}) error,
	after func(ctx context.Context, log *zap.Logger, out interface{}) error,
) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (out interface{}, err error) {
		log := GetLogger(ctx)

		if nil != before {
			err = before(ctx, log, in)
		}
		if nil != err {
			return nil, err
		}

		out, err = f(ctx, in)
		if nil != err {
			return out, err
		}

		if nil != after {
			err = after(ctx, log, out)
		}
		if nil != err {
			return nil, err
		}

		return out, nil
	}
}

func DoHTTP(
	next http.HandlerFunc,
	before func(ctx context.Context, log *zap.Logger, r *http.Request) error,
	after func(ctx context.Context, log *zap.Logger, out []byte, code int) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx := r.Context()
		log := GetLogger(ctx)

		if nil != before {
			err = before(ctx, log, r)
		}
		if nil != err {
			loghttp.Error(w, err)
			return
		}

		wr := loghttp.WrapWriter(w)

		next(wr, r)

		if nil != after {
			err = after(ctx, log, wr.Body, wr.StatusCode)
		}
		if nil != err {
			loghttp.Error(w, err)
			return
		}
	}
}
