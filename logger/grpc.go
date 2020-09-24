package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

const (
	MAX_CALLER    = 32 // reserve for max 32 caller stack
	CALLER_OFFSET = 6  // 6 default stack: runtime.goexit, serveStream, handleStream, processUnaryGRPC, grpc.handler
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
		var start time.Time
		return DoGRPC(
			f,
			func(ctx context.Context, log *zap.Logger, in interface{}) error {
				pcs := make([]uintptr, MAX_CALLER)
				n := runtime.Callers(0, pcs)

				funcName := runtime.FuncForPC(pcs[n-CALLER_OFFSET]).Name()
				log.Info(fmt.Sprintf("begin grpc request: %s", funcName))
				start = time.Now()
				return nil
			},
			func(ctx context.Context, log *zap.Logger, out interface{}) error {
				elapsed := time.Since(start)
				log.Info(fmt.Sprintf("end grpc request: %s", elapsed))
				return nil
			},
		)
	}
}

func GrpcRequestMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return DoGRPC(
			f,
			func(ctx context.Context, log *zap.Logger, in interface{}) error {
				body, _ := json.Marshal(in)
				log.Debug("grpc request payload",
					zap.String("payload", string(body)),
				)
				return nil
			},
			nil,
		)
	}
}

func GrpcResponseMiddleware() endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return DoGRPC(
			f,
			nil,
			func(ctx context.Context, log *zap.Logger, out interface{}) error {
				var by []byte
				_ = json.Unmarshal(by, out)
				log.Debug("grpc response payload",
					zap.String("payload", string(by)),
				)
				return nil
			},
		)
	}
}
