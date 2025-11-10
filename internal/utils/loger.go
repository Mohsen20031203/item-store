package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ResponseRecorder is a custom ResponseWriter that captures the status code and response body.
type ResponseRecorder struct {
	gin.ResponseWriter
	StatusCode int
	Body       []byte
}

func HTTPLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		rec := &ResponseRecorder{
			ResponseWriter: c.Writer,
			StatusCode:     http.StatusOK,
		}
		c.Writer = rec
		c.Next()
		duration := time.Since(startTime)

		logger := log.Info()
		if rec.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", rec.Body)
		}

		logger.Str("protocol", "http").
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Int("status_code", rec.StatusCode).
			Str("status_text", http.StatusText(rec.StatusCode)).
			Dur("duration", duration).
			Msg("received a HTTP request")
	}
}
