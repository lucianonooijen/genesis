package middleware

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	// MaxRequestTimeBeforeLoggingError is the amount in ms requests may take before they are too long and logged by Sentry as taking too long.
	MaxRequestTimeBeforeLoggingError = 1000
)

// TODO: Use the intrastructure logger?
// TODO: Use Sentry's ConfigureScope to set the current user in middleware
// TODO: Add custom logging fields, f.e. the LogFields

// GinLogger is Gin middleware to log request metadata and save Sentry data
func GinLogger() gin.HandlerFunc {
	var timeFormat = "02/Jan/2006:15:04:05 -0700"
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		start := time.Now()

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "request_incoming",
			Message:  fmt.Sprintf("%s request to %s", method, path),
			Level:    sentry.LevelInfo,
		})

		c.Next() // Start handling request
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		status := c.Writer.Status()

		entry := log.WithFields(log.Fields{
			"status_code": status,
			"latency":     latency,
			"client_ip":   c.ClientIP(),
			"method":      method,
			"path":        path,
			"data_length": c.Writer.Size(),
			"user_agent":  c.Request.UserAgent(),
			"timestamp":   time.Now().Format(timeFormat),
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("%d response sent for %s %s", status, method, path)
			if status > 499 {
				entry.Error(msg)
				sentry.CaptureException(errors.New(msg)) // TODO: Add traces for Sentry and create a more useful/actionable error
			} else if latency > MaxRequestTimeBeforeLoggingError {
				requestLatencyError := fmt.Sprintf("response sent, but marked as error due to request latency being too long, the latency of %dms exceeded the threshold of %dms", latency, MaxRequestTimeBeforeLoggingError)
				entry.Error(requestLatencyError)
				sentry.CaptureException(errors.New(requestLatencyError))
			} else if status > 399 {
				entry.Warn(msg)
				sentry.AddBreadcrumb(&sentry.Breadcrumb{
					Category: "response_sent",
					Message:  msg,
					Level:    sentry.LevelWarning,
				})
			} else {
				entry.Info(msg)
			}
		}
	}
}
