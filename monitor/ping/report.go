package ping

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/payfazz/fz-sentry/loghttp"
)

const (
	LEVEL_KEY     = "level"
	DEFAULT_LEVEL = 0
)

const (
	AVAILABLE                = "AVAILABLE"
	DEPENDENCY_NOT_AVAILABLE = "DEPENDENCY_NOT_AVAILABLE"
	NOT_AVAILABLE            = "NOT_AVAILABLE"
)

type ReportInterface interface {
	IsCoreService() bool
	Check(level int64) *Report
}

type Report struct {
	Service  string    `json:"service"`
	Latency  int64     `json:"latency"`
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Children []*Report `json:"children"`
	IsCore   bool      `json:"-"`
}

func GetMillisecondDuration(startRequestAt time.Time) int64 {
	return time.Since(startRequestAt).Milliseconds()
}

func Ping(serviceName string, reportChecks []ReportInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		var wg sync.WaitGroup
		children := make([]*Report, 0)
		start := time.Now()

		result := &Report{
			Service:  serviceName,
			Status:   AVAILABLE,
			Message:  "",
			Children: []*Report{},
			IsCore:   true,
		}

		level := getLevelFromQueryParam(req)

		for _, report := range reportChecks {
			wg.Add(1)
			go func(report ReportInterface) {
				defer wg.Done()
				result := report.Check(level - 1)
				if nil != result {
					children = append(children, result)
				}
			}(report)
		}

		wg.Wait()

		for _, c := range children {
			if c.Status != AVAILABLE && !c.IsCore {
				result.Status = DEPENDENCY_NOT_AVAILABLE
			}
			if c.Status != AVAILABLE && c.IsCore {
				result.Status = NOT_AVAILABLE
				break
			}
		}

		if level > 0 {
			result.Children = children
		}

		result.Latency = GetMillisecondDuration(start)

		loghttp.Write(writer, result, http.StatusOK)
	})
}

func getLevelFromQueryParam(req *http.Request) int64 {
	var level int64 = DEFAULT_LEVEL
	levels, ok := req.URL.Query()[LEVEL_KEY]
	if ok && len(levels) > 0 {
		level, _ = strconv.ParseInt(levels[0], 10, 32)
	}

	return level
}
