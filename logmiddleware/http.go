package logmiddleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/payfazz/fz-sentry/httperror"
	"github.com/payfazz/fz-sentry/logger"
	"go.uber.org/zap"
)

func DoHTTP(
	next http.HandlerFunc,
	before func(ctx context.Context, log *zap.Logger, r *http.Request) error,
	after func(ctx context.Context, log *zap.Logger, out []byte) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		ctx := r.Context()
		log := logger.GetLogger(ctx)

		err = before(ctx, log, r)
		if nil != err {
			Error(w, err)
			return
		}

		wr := WrapWriter(w)

		next(wr, r)

		err = after(ctx, log, wr.Body)
		if nil != err {
			Error(w, err)
			return
		}
	}
}

func Error(w http.ResponseWriter, err error) {
	be := httperror.GetInstance(err)
	Write(w, be, be.GetCode())
}

func Write(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

type Writer struct {
	http.ResponseWriter
	Body       []byte
	StatusCode int
}

func (w *Writer) Write(body []byte) (int, error) {
	w.Body = body
	return w.ResponseWriter.Write(body)
}

func (w *Writer) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func WrapWriter(writer http.ResponseWriter) *Writer {
	if _, ok := writer.(*Writer); ok {
		return writer.(*Writer)
	}

	return &Writer{ResponseWriter: writer}
}
