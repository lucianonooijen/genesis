package middleware

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// MaxRequestTimeBeforeLoggingError is the amount in ms requests may take before they are too long and logged by Sentry as taking too long.
	MaxRequestTimeBeforeLoggingError = 1000
)

// TODO: Add https://github.com/plimble/zap-sentry

// GinLogger is Gin middleware to log request metadata and save Sentry data.
func GinLogger(logger *zap.Logger) gin.HandlerFunc { // nolint:funlen // can't be shorter without increasing cognitive load
	timeFormat := "02/Jan/2006:15:04:05 -0700"
	httpLog := logger.Named("http_logger")

	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		start := time.Now()

		// Log the incoming request to Sentry
		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "request_incoming",
			Message:  fmt.Sprintf("%s request to %s", method, path),
			Level:    sentry.LevelInfo,
		})

		c.Next() // Start handling request

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0)) // nolint:gomnd // we have divine permission to do this
		status := c.Writer.Status()

		// Generate Zap fields
		reqData := []zap.Field{
			zap.Int("status_code", status),
			zap.Int("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("data_length", c.Writer.Size()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("timestamp", time.Now().Format(timeFormat)),
		}

		// Log breadcrumb to Sentry
		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Level: sentry.LevelInfo,
			Data: map[string]interface{}{
				"status_code": status,
				"latency":     latency,
				"client_ip":   c.ClientIP(),
				"method":      method,
				"path":        path,
				"data_length": c.Writer.Size(),
				"user_agent":  c.Request.UserAgent(),
				"timestamp":   time.Now().Format(timeFormat),
			},
		})

		// Log if there are errors in the Gin context
		if len(c.Errors) > 0 {
			err := c.Errors.ByType(gin.ErrorTypePrivate).String()
			httpLog.Error(err, reqData...)
			sentry.CaptureException(errors.New(err)) // TODO: Maybe remove?

			return
		}

		msg := fmt.Sprintf("%d response sent for %s %s", status, method, path)

		// Log server errors
		if status > 499 { // nolint:gomnd // use for http code limits are ok
			httpLog.Error(msg, reqData...)
			sentry.CaptureException(errors.New(msg)) // TODO: Add traces for Sentry and create a more useful/actionable error

			return
		}

		// Log requests that take too long
		if latency > MaxRequestTimeBeforeLoggingError {
			requestLatencyError := fmt.Sprintf("response sent, but marked as error due to request latency being too long, the latency of %dms exceeded the threshold of %dms", latency, MaxRequestTimeBeforeLoggingError)
			httpLog.Warn(requestLatencyError, reqData...)
			sentry.CaptureException(errors.New(requestLatencyError))

			return
		}

		// Log 4xx responses as warnings
		// Check for < 500 not needed because of the > 499 check and return above
		if status > 399 { // nolint:gomnd // use for http code limits are ok
			httpLog.Warn(msg, reqData...)
			sentry.AddBreadcrumb(&sentry.Breadcrumb{
				Category: "response_sent",
				Message:  msg,
				Level:    sentry.LevelWarning,
			})

			return
		}

		httpLog.Info(msg, reqData...)
	}
}
