package middleware

import (
	"net/http"
	"time"

	router "github.com/payfazz/fz-router"
	"github.com/payfazz/fz-sentry/loghttp"
	"github.com/payfazz/fz-sentry/monitor/prometheusclient"
)

// HTTPRequestCounterMiddleware middleware wrapper for IncrementRequestCounter, recommended to be used if you are using `go-apt/pkg/fazzrouter` package, the only thing required: before using this middleware make sure you use `kv.New()` middleware from `github.com/payfazz/go-middleware`
func HTTPRequestCounterMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, req *http.Request) {
			prometheusWriter := loghttp.WrapWriter(writer)

			next(prometheusWriter, req)

			prometheusclient.IncrementRequestCounter(
				router.GetPattern(req),
				req.Method,
				prometheusWriter.Code(),
			)
		}
	}
}

// HTTPRequestDurationMiddleware middleware wrapper for ObserveRequestDuration, recommended to be used if you are using `go-apt/pkg/fazzrouter` package, the only thing required: before using this middleware make sure you use `kv.New()` middleware from `github.com/payfazz/go-middleware`
func HTTPRequestDurationMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, req *http.Request) {
			start := time.Now()
			prometheusWriter := loghttp.WrapWriter(writer)

			next(prometheusWriter, req)

			prometheusclient.ObserveRequestDuration(
				router.GetPattern(req),
				req.Method,
				prometheusWriter.Code(),
				start,
			)
		}
	}
}
