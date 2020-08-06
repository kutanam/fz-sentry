package logger

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func HttpMiddleware(logger *zap.Logger) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = NewRequest(ctx, logger)
			next(w, r.WithContext(ctx))
		}
	}
}

func HttpEndpointMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		var start time.Time
		return DoHTTP(
			next,
			func(ctx context.Context, log *zap.Logger, r *http.Request) error {
				log.Info(fmt.Sprintf("begin http request: %-15s | %-10s | %-15s | %-40s", getIP(r), r.Method, r.Host, r.RequestURI))
				start = time.Now()
				return nil
			},
			func(ctx context.Context, log *zap.Logger, out []byte, code int) error {
				elapsed := time.Since(start)
				log.Info(fmt.Sprintf("end http request: %s", elapsed))
				return nil
			},
		)
	}
}

func HttpRequestMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return DoHTTP(
			next,
			func(ctx context.Context, log *zap.Logger, r *http.Request) error {
				buf, _ := ioutil.ReadAll(r.Body)
				body := ioutil.NopCloser(bytes.NewBuffer(buf))
				r.Body = body
				log.Debug("http request payload",
					zap.String("payload", string(buf)),
				)
				return nil
			},
			nil,
		)
	}
}

func HttpResponseMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return DoHTTP(
			next,
			nil,
			func(ctx context.Context, log *zap.Logger, out []byte, code int) error {
				log.Debug("http response payload",
					zap.String("payload", string(out)),
					zap.Int("http status", code),
				)
				return nil
			},
		)
	}
}
