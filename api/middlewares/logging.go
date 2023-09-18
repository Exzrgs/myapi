package middlewares

import (
	"log"
	"net/http"

	"github.com/Exzrgs/myapi/common"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func NewLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func (w *resLoggingWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := common.SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)
		lw := NewLoggingWriter(w)

		next.ServeHTTP(lw, req)
		log.Printf("[%d]res: %d\n", traceID, lw.code)
	})
}
