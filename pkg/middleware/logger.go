package middleware

import (
	"bytes"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type ResponseRecorder struct {
	gin.ResponseWriter
	StatusCode int
	body       *bytes.Buffer
}

func (w *ResponseRecorder) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w ResponseRecorder) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		log := zerolog.New(output).With().Timestamp().Logger()
		start := time.Now()

		blw := &ResponseRecorder{ResponseWriter: c.Writer, StatusCode: http.StatusOK, body: bytes.NewBufferString("")}
		c.Writer = blw
		c.Next()
		logger := log.Info()
		if c.Writer.Status() != http.StatusOK {
			logger = log.Error().Bytes("body", blw.body.Bytes())
		}
		duration := time.Since(start)
		logger.Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("remote_addr", c.Request.RemoteAddr).
			Str("user_agent", c.Request.UserAgent()).
			Str("time", start.Format(time.RFC3339)).
			Int("status_code", c.Writer.Status()).
			Str("status_text", http.StatusText(c.Writer.Status())).
			Dur("duration", duration).
			Dur("body", duration).
			Msg("[GIN API Server]")
	}
}

// func HttpLogger(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		rec := &ResponseRecorder{ResponseWriter: w, StatusCode: http.StatusOK}

// 		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
// 		next.ServeHTTP(rec, r)
// 		duration := time.Since(start)
// 		logger := log.Info()
// 		if rec.StatusCode != http.StatusOK {
// 			logger = log.Error().Bytes("body", rec.Body)
// 		}
// 		logger.Str("method", r.Method).
// 			Str("path", r.URL.Path).
// 			Str("remote_addr", r.RemoteAddr).
// 			Str("user_agent", r.UserAgent()).
// 			Str("time", start.Format(time.RFC3339)).
// 			Int("status_code", rec.StatusCode).
// 			Str("status_text", http.StatusText(rec.StatusCode)).
// 			Dur("duration", duration).
// 			Msg("request via HTTP")
// 	})
// }
