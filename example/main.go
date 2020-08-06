package main

import (
	"net/http"

	"github.com/payfazz/fz-sentry/example/controller"
	"github.com/payfazz/fz-sentry/logger"
)

const (
	ENV_PRODUCTION  = "production"
	ENV_DEVELOPMENT = "development"
)

func wrapMiddleware(next http.HandlerFunc, env string) http.HandlerFunc {
	log := logger.New(env, "example")

	return logger.HttpMiddleware(log)(logger.HttpResponseMiddleware()(next))
}

func main() {
	devServer()
	prodServer()

	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		panic(err)
	}
}

func devServer() {
	http.Handle("/dev/success", wrapMiddleware(controller.Success(), ENV_DEVELOPMENT))
	/*
		2020-07-29T18:08:27.290+0700    DEBUG   controller/success.go:15        this is debug message   {"serviceName": "example", "requestId": "e8469305-580c-4feb-a5a3-f8fa9a4f5643"}
		2020-07-29T18:08:27.290+0700    INFO    controller/success.go:17        this is info from success controller    {"serviceName": "example", "requestId": "e8469305-580c-4feb-a5a3-f8fa9a4f5643", "status": "request processed successfully"}
	*/

	http.Handle("/dev/warning", wrapMiddleware(controller.Warning(), ENV_DEVELOPMENT))
	/*
		2020-07-29T18:08:22.866+0700    DEBUG   controller/warning.go:15        this is debug message   {"serviceName": "example", "requestId": "44adbf96-1b48-4817-9178-6774fd2b2826"}
		2020-07-29T18:08:22.866+0700    WARN    controller/warning.go:17        this is warning {"serviceName": "example", "requestId": "44adbf96-1b48-4817-9178-6774fd2b2826", "cause": "warning occured"}
		github.com/payfazz/fz-sentry/example/controller.Warning.func1
		        /Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/controller/warning.go:17
		github.com/payfazz/fz-sentry/example/middleware.Logger.func1
		        /Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/middleware/logger.go:19
		net/http.HandlerFunc.ServeHTTP
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2007
		net/http.(*ServeMux).ServeHTTP
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2387
		net/http.serverHandler.ServeHTTP
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2802
		net/http.(*conn).serve
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:1890
	*/

	http.Handle("/dev/error", wrapMiddleware(controller.Error(), ENV_DEVELOPMENT))
	/*
		2020-07-29T18:07:59.998+0700    DEBUG   controller/error.go:16  this is debug message   {"serviceName": "example", "requestId": "fdd65fa9-8838-4249-8985-733bc10a1a27"}
		2020-07-29T18:07:59.998+0700    ERROR   controller/error.go:19  this is error   {"serviceName": "example", "requestId": "fdd65fa9-8838-4249-8985-733bc10a1a27", "cause": "undefined error"}
		github.com/payfazz/fz-sentry/example/controller.Error.func1
		        /Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/controller/error.go:19
		github.com/payfazz/fz-sentry/example/middleware.Logger.func1
		        /Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/middleware/logger.go:19
		net/http.HandlerFunc.ServeHTTP
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2007
		net/http.(*ServeMux).ServeHTTP
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2387
		net/http.serverHandler.ServeHTTP
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2802
		net/http.(*conn).serve
		        /usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:1890
	*/
}

func prodServer() {
	http.Handle("/prd/success", wrapMiddleware(controller.Success(), ENV_PRODUCTION))
	// {"level":"info","ts":"2020-07-30T01:28:11+07:00","caller":"controller/success.go:17","msg":"this is info from success controller","serviceName":"example","requestId":"043358ba-5c49-4019-bfd5-cc62abed5279","status":"request processed successfully"}

	http.Handle("/prd/warning", wrapMiddleware(controller.Warning(), ENV_PRODUCTION))
	// {"level":"warn","ts":"2020-07-30T01:28:11+07:00","caller":"controller/warning.go:17","msg":"this is warning","serviceName":"example","requestId":"5a03a1e0-5443-41c8-b3b1-749094532147","cause":"warning occured"}

	http.Handle("/prd/error", wrapMiddleware(controller.Error(), ENV_PRODUCTION))
	// {"level":"error","ts":"2020-07-30T01:28:11+07:00","caller":"controller/error.go:19","msg":"this is error","serviceName":"example","requestId":"226d6478-9a30-47fa-adbc-30386bb65383","cause":"undefined error","stacktrace":"github.com/payfazz/fz-sentry/example/controller.Error.func1\n\t/Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/controller/error.go:19\ngithub.com/payfazz/fz-sentry/example/middleware.Logger.func1\n\t/Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/middleware/logger.go:19\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2007\nnet/http.(*ServeMux).ServeHTTP\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2387\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2802\nnet/http.(*conn).serve\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:1890"}
}
