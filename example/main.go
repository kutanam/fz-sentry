package main

import (
	"net/http"

	"github.com/payfazz/fz-sentry/example/controller"
	"github.com/payfazz/fz-sentry/example/middleware"
)

const (
	ENV_PRODUCTION  = "production"
	ENV_DEVELOPMENT = "development"
)

func main() {
	devServer()
	prodServer()

	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		panic(err)
	}
}

func devServer() {
	http.Handle("/dev/success", middleware.Logger(controller.Success(), ENV_DEVELOPMENT, ""))
	/*
		2020-07-29T18:08:27.290+0700    DEBUG   controller/success.go:15        this is debug message   {"serviceName": "example", "requestId": "e8469305-580c-4feb-a5a3-f8fa9a4f5643"}
		2020-07-29T18:08:27.290+0700    INFO    controller/success.go:17        this is info from success controller    {"serviceName": "example", "requestId": "e8469305-580c-4feb-a5a3-f8fa9a4f5643", "status": "request processed successfully"}
	*/

	http.Handle("/dev/warning", middleware.Logger(controller.Warning(), ENV_DEVELOPMENT, ""))
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

	http.Handle("/dev/error", middleware.Logger(controller.Error(), ENV_DEVELOPMENT, ""))
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
	http.Handle("/prd/success", middleware.Logger(controller.Success(), ENV_PRODUCTION, ""))
	// {"level":"info","ts":1596020749.78262,"caller":"controller/success.go:17","msg":"this is info from success controller","serviceName":"example","requestId":"043358ba-5c49-4019-bfd5-cc62abed5279","status":"request processed successfully"}

	http.Handle("/prd/warning", middleware.Logger(controller.Warning(), ENV_PRODUCTION, ""))
	// {"level":"warn","ts":1596020742.776434,"caller":"controller/warning.go:17","msg":"this is warning","serviceName":"example","requestId":"5a03a1e0-5443-41c8-b3b1-749094532147","cause":"warning occured"}

	http.Handle("/prd/error", middleware.Logger(controller.Error(), ENV_PRODUCTION, ""))
	// {"level":"error","ts":1596021222.586966,"caller":"controller/error.go:19","msg":"this is error","serviceName":"example","requestId":"16866cb0-2caa-4c89-8315-7d3df1a35462","cause":"undefined error","stacktrace":"github.com/payfazz/fz-sentry/example/controller.Error.func1\n\t/Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/controller/error.go:19\ngithub.com/payfazz/fz-sentry/example/middleware.Logger.func1\n\t/Users/cashfazz002/go/src/github.com/payfazz/fz-sentry/example/middleware/logger.go:19\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2007\nnet/http.(*ServeMux).ServeHTTP\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2387\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:2802\nnet/http.(*conn).serve\n\t/usr/local/Cellar/go/1.13.3/libexec/src/net/http/server.go:1890"}
}