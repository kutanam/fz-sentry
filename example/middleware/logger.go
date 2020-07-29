package middleware

import (
	"net/http"

	"github.com/payfazz/fz-sentry/logger"
)

func Logger(next http.HandlerFunc, env string, slackHookUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.New(env, "example")
		if "" != slackHookUrl {
			log = logger.New(env, "example", logger.DebugSlackHook(slackHookUrl))
		}

		ctx = logger.NewRequest(ctx, log)

		next(w, r.WithContext(ctx))
	}
}
