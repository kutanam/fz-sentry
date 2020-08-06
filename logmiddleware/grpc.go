package logmiddleware

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/payfazz/fz-sentry/logger"
	"go.uber.org/zap"
)

func DoGRPC(
	f endpoint.Endpoint,
	before func(ctx context.Context, log *zap.Logger, in interface{}) error,
	after func(ctx context.Context, log *zap.Logger, out interface{}) error,
) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (out interface{}, err error) {
		log := logger.GetLogger(ctx)

		err = before(ctx, log, in)
		if nil != err {
			return nil, err
		}

		out, err = f(ctx, in)
		if nil != err {
			return out, err
		}

		err = after(ctx, log, out)
		if nil != err {
			return nil, err
		}

		return out, nil
	}
}
