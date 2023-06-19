package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

// DefaultStructuredLogger logs a gin HTTP request in JSON format
func DefaultStructuredLogger() gin.HandlerFunc {
	return StructuredLogger(&log.Logger)
}

// StructuredLogger logs an HTTP request in JSON format.
func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {

	setMinimumLogLevelForGinMode()

	return func(c *gin.Context) {

		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := getLogParametersForCurrentContext(c, start, raw, path)

		logEvent := setupLogEventForCurrentContext(c, logger)
		setLogEventFieldsFromLogFormatter(logEvent, param)
		logEvent.Msg(param.ErrorMessage)
	}
}

func setMinimumLogLevelForGinMode() {
	isReleaseMode := os.Getenv("GIN_MODE") == "release"

	if isReleaseMode {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func getLogParametersForCurrentContext(c *gin.Context, start time.Time, raw string, path string) gin.LogFormatterParams {
	param := gin.LogFormatterParams{}

	param.TimeStamp = time.Now() // Stop timer
	param.Latency = param.TimeStamp.Sub(start)
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	param.ClientIP = c.ClientIP()
	param.Method = c.Request.Method
	param.StatusCode = c.Writer.Status()
	param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
	param.BodySize = c.Writer.Size()
	if raw != "" {
		path = path + "?" + raw
	}
	param.Path = path
	return param
}

func setupLogEventForCurrentContext(c *gin.Context, logger *zerolog.Logger) *zerolog.Event {
	var logEvent *zerolog.Event
	if c.Writer.Status() >= 500 {
		logEvent = logger.Error()
	} else {
		if strings.Contains(c.Request.URL.Path, "healthcheck") {
			logEvent = logger.Debug()
		} else {
			logEvent = logger.Info()
		}
	}

	return logEvent
}

func setLogEventFieldsFromLogFormatter(logEvent *zerolog.Event, param gin.LogFormatterParams) {
	logEvent.
		Str("client_id", param.ClientIP).
		Str("method", param.Method).
		Int("status_code", param.StatusCode).
		Int("body_size", param.BodySize).
		Str("path", param.Path).
		Str("latency", param.Latency.String())
}
