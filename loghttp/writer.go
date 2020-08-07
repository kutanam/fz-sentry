package loghttp

import (
	"fmt"
	"net/http"
)

type Writer struct {
	http.ResponseWriter
	Body       []byte
	StatusCode int
}

func (w *Writer) Code() string {
	return fmt.Sprint(w.StatusCode)
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
