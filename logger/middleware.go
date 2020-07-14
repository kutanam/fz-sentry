package logger

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

func HttpMiddleware(logger *zap.Logger) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = NewLoggerContext(ctx, logger)
			next(w, r)
		}
	}
}

func HttpEndpointMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := GetLogger(ctx)

			logger.Info(fmt.Sprintf("begin http request: %-15s | %-10s | %-15s | %-40s", getIP(r), r.Method, r.Host, r.RequestURI))

			start := time.Now()
			next(w, r)
			elapsed := time.Since(start)

			logger.Info(fmt.Sprintf("end http request: %s", elapsed))
		}
	}
}

func HttpRequestMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := GetLogger(ctx)

			buf, _ := ioutil.ReadAll(r.Body)
			body := ioutil.NopCloser(bytes.NewBuffer(buf))
			r.Body = body

			logger.Debug("http request payload",
				zap.String("payload", string(buf)),
			)

			next(w, r)
		}
	}
}

func GrpcMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			ctx = NewLoggerContext(ctx, logger)
			return f(ctx, in)
		}
	}
}
