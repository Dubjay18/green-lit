// Package logger provides a simple logging utility using the zap library.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log is the global logger used by the application.
var log *zap.Logger

// init initializes the global logger with a production configuration.
func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"                   // Set the key for the timestamp field.
	encoderConfig.StacktraceKey = ""                      // Disable the stacktrace field.
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO8601 format for timestamps.
	config.EncoderConfig = encoderConfig

	// Build the logger with the configured settings.
	// zap.AddCallerSkip(1) skips one caller frame when annotating logs with file and line number information.
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err) // Panic if the logger cannot be built.
	}
}

// Info logs an informational message with optional structured context fields.
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug logs a debug message with optional structured context fields.
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error logs an error message with optional structured context fields.
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
