// Package log defines common logging interface with basic implementation.
package log

import "os"

// Logger defines common logger interface.
type Logger interface {
	// Debug log.
	Debug(msg string, args ...interface{})
	// Info log.
	Info(msg string, args ...interface{})
	// Warning log.
	Warning(msg string, args ...interface{})
	// Error log.
	Error(msg string, args ...interface{})
	// Fatal log.
	Fatal(msg string, args ...interface{})
}

// Default logger.
var logger = defaultLogger

// GetDefaultLogger returns builtin Logger implementation
func GetDefaultLogger() Logger {
	return defaultLogger
}

// GetLogger returns current default logger
func GetLogger() Logger {
	return logger
}

// SetLogger replaces default Logger.
func SetLogger(l Logger) {
	logger = l
}

// Debug log.
func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

// Info log.
func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

// Warning log.
func Warning(msg string, args ...interface{}) {
	logger.Warning(msg, args...)
}

// Error log.
func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

// Fatal log.
func Fatal(msg string, args ...interface{}) {
	logger.Fatal(msg, args...)
	os.Exit(1)
}
