// Package log defines common logging interface with basic implementation.
package log

import "os"

// TODO consider to use github.com/golang/glog for default logger
import log "github.com/op/go-logging"

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

// Global logger.
var logger = defaultLogger()

// SetLogger replaces global Logger.
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

// Logger impl.
type loggerImpl struct {
	log *log.Logger
}

// Creates default logger.
func defaultLogger() Logger {
	var module = os.Args[0]
	var logger = log.MustGetLogger(module)
	var format = "[" + module + "] %{color}%{time:15:04:05.000000} %{message}%{color:reset}"
	log.SetFormatter(log.MustStringFormatter(format))

	if level := os.Getenv("LOG_LEVEL"); level != "" {
		var levels = map[string]log.Level{
			"debug":    log.DEBUG,
			"info":     log.INFO,
			"notice":   log.NOTICE,
			"warning":  log.WARNING,
			"error":    log.ERROR,
			"critical": log.CRITICAL,
			"fatal":    log.CRITICAL,
		}
		log.SetLevel(levels[level], module)
	} else {
		log.SetLevel(log.DEBUG, module)
	}

	return &loggerImpl{logger}
}

// Debug log.
func (l *loggerImpl) Debug(msg string, args ...interface{}) {
	l.log.Debug(msg, args)
}

// Info log.
func (l *loggerImpl) Info(msg string, args ...interface{}) {
	l.log.Info(msg, args)
}

// Warning log.
func (l *loggerImpl) Warning(msg string, args ...interface{}) {
	l.log.Warning(msg, args)
}

// Error log.
func (l *loggerImpl) Error(msg string, args ...interface{}) {
	l.log.Error(msg, args)
}

// Fatal log.
func (l *loggerImpl) Fatal(msg string, args ...interface{}) {
	l.log.Fatalf(msg, args)
}
