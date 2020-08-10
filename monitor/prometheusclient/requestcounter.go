package prometheusclient

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var counterOnce sync.Once
var counter *prometheus.CounterVec

func requestCounter() *prometheus.CounterVec {
	counterOnce.Do(func() {
		counter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "_requests_total",
				Help: "A counter for requests to the wrapped handler.",
			},
			[]string{"path", "method", "code"},
		)

		prometheus.MustRegister(counter)
	})

	return counter
}

// IncrementRequestCounter increment  request count and store it as total requests, usage example can be seen in HTTPRequestCounterMiddleware method
// required params:
// - serviceName: your service name (snake_case)
// - pattern: your route pattern not the requested url, ex: `/v1/users/:id` (correct); `/v1/users/{id}` (correct); `/v1/users/1` (incorrect)
// - method: your  request method (GET, POST, PATCH, etc)
// - code: your  status code (200, 404, 500, etc)
func IncrementRequestCounter(pattern string, method string, code string) {
	labels := prometheus.Labels{
		"path":   pattern,
		"method": method,
		"code":   code,
	}
	requestCounter().With(labels).Inc()
}
