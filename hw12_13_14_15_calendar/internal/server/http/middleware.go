package internalhttp

import (
	"net/http"
	"strings"
	"time"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK, 0}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.size += size
	return size, err
}

func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)

		tstart := time.Now()
		next.ServeHTTP(lrw, r)
		code := lrw.statusCode
		size := lrw.size
		timeToTakeServ := time.Since(tstart)
		raddrlog := r.RemoteAddr[0:strings.Index(r.RemoteAddr, ":")]
		timelog := tstart.UTC().Format("02/01/2006 15:04:05 UTC")

		s.logHTTP.httplogger.Printf("%s [%s] %s %s %s %d %d %s %s\n",
			raddrlog, timelog, r.Method, r.URL.Path, r.Proto, code, size, timeToTakeServ, r.UserAgent())
	})
}
