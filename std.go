package log

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	debugLevel int = iota
	infoLevel
	warnLevel
	errorLevel
	criticalLevel
)

var (
	colors = []*color.Color{
		debugLevel:    color.New(color.FgCyan),
		infoLevel:     color.New(color.FgWhite),
		warnLevel:     color.New(color.FgYellow),
		errorLevel:    color.New(color.FgRed),
		criticalLevel: color.New(color.FgMagenta),
	}
	defaultLogger = makeDefaultLogger()
)

// Logger impl.
type loggerImpl struct {
	prefix string
	level  int
	sync.Mutex
}

// Creates default logger.
func makeDefaultLogger() Logger {
	var level = debugLevel

	if val := os.Getenv("LOG_LEVEL"); val != "" {
		var levels = map[string]int{
			"debug":    debugLevel,
			"info":     infoLevel,
			"warning":  warnLevel,
			"warn":     warnLevel,
			"error":    errorLevel,
			"critical": criticalLevel,
			"fatal":    criticalLevel,
		}
		level = levels[val]
	}

	return &loggerImpl{level: level}
}

// Debug log.
func (l *loggerImpl) Debug(msg string, args ...interface{}) {
	if l.level > debugLevel {
		return
	}
	l.print(debugLevel, "[DBG]", msg, args...)
}

// Info log.
func (l *loggerImpl) Info(msg string, args ...interface{}) {
	if l.level > infoLevel {
		return
	}
	l.print(infoLevel, "[INF]", msg, args...)
}

// Warning log.
func (l *loggerImpl) Warning(msg string, args ...interface{}) {
	if l.level > warnLevel {
		return
	}
	l.print(warnLevel, "[WRN]", msg, args...)
}

// Error log.
func (l *loggerImpl) Error(msg string, args ...interface{}) {
	if l.level > errorLevel {
		return
	}
	l.print(errorLevel, "[ERR]", msg, args...)
}

// Fatal log.
func (l *loggerImpl) Fatal(msg string, args ...interface{}) {
	if l.level > criticalLevel {
		return
	}
	l.print(criticalLevel, "[FAT]", msg, args...)
}

// Print a message.
func (l *loggerImpl) print(level int, lstr string, msg string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()
	var t = time.Now().Format("15:04:05.000000")
	var c = colors[level]
	c.Println(fmt.Sprintf("%s%s %s %s", l.prefix, lstr, t, fmt.Sprintf(msg, args...)))
}
