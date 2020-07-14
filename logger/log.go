package logger

import (
	"context"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type loggerKeyType struct{}

var loggerKey loggerKeyType

func GetLogger(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey).(*zap.Logger)
}

func NewLoggerContext(ctx context.Context, logger *zap.Logger) context.Context {
	requestId, _ := uuid.NewV4()
	logger = logger.With(
		zap.String("requestId", requestId.String()),
	)
	return context.WithValue(ctx, loggerKey, logger)
}

func New(env string, options ...zap.Option) *zap.Logger {
	logger, _ := zap.NewDevelopment(options...)
	if "development" != env {
		logger, _ = zap.NewProduction(options...)
	}

	return logger
}
