package logger

import (
	"time"

	"git.bytecode.nl/bytecode/genesis/internal/interactors"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

/*
 * A small note on this package:
 * This package is meant as a simple wrapper to avoid importing the logrus
 * library in the code directly, and also to log errors and such to Sentry
 * without needing a separate call in the production code. Note that for
 * specific logging, like incoming requests, the logrus package is used directly
 */

// Logger contains methods for logging errors, warnings, info and debug messages.
//
// The logging settings must be set using Configure on application init, before using Logger.
//
// All methods will send data both to logrus and Sentry.
type Logger struct {
	location string
}

// TODO: Add wrapper method for logging thrown errors, f.e. `return LogIfError(err)`, so that when developing we can always trace created errors everywhere
// TODO: Add functionality for tracing errors/function calls (like automatic `trace` logging) and/or stacktrace like support (the 'history' of called functions) when debugging locally
// TODO: Clean up duplicate code

// Configure configures logrus and Sentry. Should only be called once when initting the application.
func Configure(isDevMode bool, sentryDSN string, sentryEnv string) error {
	if isDevMode {
		log.SetLevel(log.TraceLevel) // Lowest level in Logrus.
	} else { // Is production
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         sentryDSN,
			Environment: sentryEnv,
		}); err != nil {
			return err
		}
		sentry.Flush(time.Second * 5)
	}
	return nil
}

// New returns a Logger instance
func New(location string) interactors.Logger {
	return Logger{
		location: location,
	}
}

func (l Logger) saveAndGetFields(additionalFields interactors.LogFields) log.Fields {
	additionalFields["location"] = l.location
	// TODO: Also add runtime information, from where the error was called, example
	//_, file, line, _ := runtime.Caller(3) // Must be 3/4 depending on X or XWithFields being called
	//additionalFields["called_from"] = file + ":" + strconv.Itoa(line)
	fields := log.Fields(additionalFields)
	sentry.AddBreadcrumb(&sentry.Breadcrumb{Data: fields})
	return fields
}

// Fatal logs an error as fatal, exits with exit code 1
// Should only be used in the main function
func (l Logger) Fatal(err error) {
	sentry.CaptureException(err)
	logWithFields := log.WithFields(l.saveAndGetFields(interactors.LogFields{}))
	logWithFields.Fatal(err.Error())
}

// Error logs an error as error
// To pass additional information use ErrorWithFields
func (l Logger) Error(err error) {
	l.ErrorWithFields(err, interactors.LogFields{})
}

// ErrorWithFields logs an an error with extra LogFields as error.
// To call without LogFields, use Error
func (l Logger) ErrorWithFields(err error, fields interactors.LogFields) {
	logWithFields := log.WithFields(l.saveAndGetFields(fields))
	logWithFields.Error(err.Error())
	sentry.CaptureException(err)
}

// Warn logs an error as warning
// To pass additional information use WarnWithFields
func (l Logger) Warn(err error) {
	l.WarnWithFields(err, interactors.LogFields{})
}

// WarnWithFields logs an an error with extra LogFields as warning.
// To call without LogFields, use Warn
func (l Logger) WarnWithFields(err error, fields interactors.LogFields) {
	logWithFields := log.WithFields(l.saveAndGetFields(fields))
	logWithFields.Warn(err.Error())
	sentry.CaptureMessage(err.Error())
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Message: err.Error(),
		Level:   sentry.LevelWarning,
	})
}

// Info logs a string as info.
// To pass additional information use InfoWithFields
func (l Logger) Info(message string) {
	l.InfoWithFields(message, interactors.LogFields{})
}

// InfoWithFields logs a string with extra LogFields as info.
// To call without LogFields, use Info
func (l Logger) InfoWithFields(message string, fields interactors.LogFields) {
	logWithFields := log.WithFields(l.saveAndGetFields(fields))
	logWithFields.Info(message)
	sentry.CaptureMessage(message)
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Message: message,
		Level:   sentry.LevelInfo,
	})
}

// Debug logs a string with debug severity.
// To pass additional information use DebugWithFields
func (l Logger) Debug(message string) {
	l.DebugWithFields(message, interactors.LogFields{})
}

// DebugWithFields logs a string with debug severity.
// To call without LogFields, use Debug
func (l Logger) DebugWithFields(message string, fields interactors.LogFields) {
	logWithFields := log.WithFields(l.saveAndGetFields(fields))
	logWithFields.Debug(message)
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Message: message,
		Level:   sentry.LevelDebug,
	})
}

// Trace logs a string with debug severity.
// To pass additional information use TraceWithFields
func (l Logger) Trace(message string) {
	l.TraceWithFields(message, interactors.LogFields{})
}

// TraceWithFields logs a string with trace severity.
// To call without LogFields, use Trace
func (l Logger) TraceWithFields(message string, fields interactors.LogFields) {
	logWithFields := log.WithFields(l.saveAndGetFields(fields))
	logWithFields.Trace(message)
	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Message: message,
		Level:   sentry.LevelDebug,
	})
}
