package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var configDev = zap.Config{
	Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
	Development: true,
	Encoding:    "console",
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "name",
		CallerKey:        "caller",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("15:04:05.000"),
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "\t",
	},
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
}

var configProd = zap.Config{
	Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
	Development: false,
	Encoding:    "json",
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "name",
		CallerKey:        "caller",
		FunctionKey:      "function",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "\t",
	},
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
	InitialFields:    nil,
}

// Config returns the correct logging configuration for the given environment
func Config(isDev bool) zap.Config {
	// DEVELOPMENT
	if isDev {
		return configDev
	}

	// PRODUCTION
	return configProd
}

// New returns a *zap.Logger for the given environment
func New(isDev bool) (*zap.Logger, error) {
	config := Config(isDev)
	logger, err := config.Build()
	return logger, err
}
