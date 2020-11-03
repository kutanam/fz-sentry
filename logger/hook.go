package logger

import (
	"github.com/payfazz/fz-sentry/slackcore"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SlackHook(slackHookUrl string, outLevel zapcore.Level) zap.Option {
	return zap.WrapCore(
		slackcore.NewWrapper(
			slackHookUrl,
			outLevel,
		),
	)
}

func DebugSlackHook(slackHookUrl string) zap.Option {
	return SlackHook(slackHookUrl, zapcore.DebugLevel)
}

func InfoSlackHook(slackHookUrl string) zap.Option {
	return SlackHook(slackHookUrl, zapcore.InfoLevel)
}

func WarnSlackHook(slackHookUrl string) zap.Option {
	return SlackHook(slackHookUrl, zapcore.WarnLevel)
}

func ErrorSlackHook(slackHookUrl string) zap.Option {
	return SlackHook(slackHookUrl, zapcore.ErrorLevel)
}
