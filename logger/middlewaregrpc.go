package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

func GrpcMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			ctx = NewRequest(ctx, logger)
			return f(ctx, in)
		}
	}
}

func GrpcEndpointMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			logger := GetLogger(ctx)

			funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			logger.Info(fmt.Sprintf("begin grpc request: %s", funcName))

			start := time.Now()
			out, err = f(ctx, in)
			elapsed := time.Since(start)

			logger.Info(fmt.Sprintf("end grpc request: %s", elapsed))
			return out, err
		}
	}
}

func GrpcRequestMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			logger := GetLogger(ctx)
			body, _ := json.Marshal(in)

			logger.Debug("grpc request payload",
				zap.String("payload", string(body)),
			)

			return f(ctx, in)
		}
	}
}

func GrpcResponseMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, in interface{}) (out interface{}, err error) {
			logger := GetLogger(ctx)

			out, err = f(ctx, in)

			var by []byte
			_ = json.Unmarshal(by, out)

			logger.Debug("grpc response payload",
				zap.String("payload", string(by)),
			)

			return out, err
		}
	}
}
