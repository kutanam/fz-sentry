package prometheusclient

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/go-kit/kit/endpoint"
	router "github.com/payfazz/fz-router"
	"github.com/payfazz/fz-sentry/httperror"
	"github.com/payfazz/fz-sentry/loghttp"
)

const (
	GRPC = "grpc"
)

// HTTPRequestCounterMiddleware middleware wrapper for IncrementRequestCounter, recommended to be used if you are using `fz-router` package
func HTTPRequestCounterMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, req *http.Request) {
			prometheusWriter := loghttp.WrapWriter(writer)

			next(prometheusWriter, req)

			IncrementRequestCounter(
				router.GetPattern(req),
				req.Method,
				prometheusWriter.Code(),
			)
		}
	}
}

// HTTPRequestDurationMiddleware middleware wrapper for ObserveRequestDuration, recommended to be used if you are using `fz-router` package
func HTTPRequestDurationMiddleware() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, req *http.Request) {
			start := time.Now()
			prometheusWriter := loghttp.WrapWriter(writer)

			next(prometheusWriter, req)

			ObserveRequestDuration(
				router.GetPattern(req),
				req.Method,
				prometheusWriter.Code(),
				start,
			)
		}
	}
}

// GRPCRequestCounterMiddleware middleware wrapper for IncrementRequestCounter, recommended to be used if you are using `fz-router` package
func GRPCRequestCounterMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			out, err = f(ctx, in)

			code := 200
			if nil != err {
				be := httperror.GetInstance(err)
				code = be.GetCode()
			}

			IncrementRequestCounter(
				runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
				GRPC,
				fmt.Sprint(code),
			)

			return out, err
		}
	}
}

// GRPCRequestDurationMiddleware middleware wrapper for ObserveRequestDuration, recommended to be used if you are using `fz-router` package
func GRPCRequestDurationMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			start := time.Now()

			out, err = f(ctx, in)

			code := 200
			if nil != err {
				be := httperror.GetInstance(err)
				code = be.GetCode()
			}

			ObserveRequestDuration(
				runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
				GRPC,
				fmt.Sprint(code),
				start,
			)

			return out, err
		}
	}
}
