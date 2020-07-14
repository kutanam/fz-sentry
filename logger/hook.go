package logger

import (
	"github.com/bluele/zapslack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SlackHook(slackHookUrl string, outLevel zapcore.Level) zap.Option {
	return zap.Hooks(zapslack.NewSlackHook(slackHookUrl, outLevel).GetHook())
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
