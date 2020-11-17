package interactors

// LogFields is the type for optional arguments when logging
type LogFields map[string]interface{}

// Logger is the interface used for logging throughout the application
type Logger interface {
	Error(err error)
	ErrorWithFields(err error, fields LogFields)
	Warn(err error)
	WarnWithFields(err error, fields LogFields)
	Info(message string)
	InfoWithFields(message string, fields LogFields)
	Debug(message string)
	DebugWithFields(message string, fields LogFields)
}
