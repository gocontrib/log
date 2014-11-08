package log

import "fmt"
import "os"
import "sync"
import "time"
import "github.com/ttacon/chalk"

const (
	DEBUG int = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

// Logger impl.
type loggerImpl struct {
	module string
	level  int
	sync.Mutex
}

// Creates default logger.
func defaultLogger() Logger {
	var module = os.Args[0]
	var level = DEBUG

	if val := os.Getenv("LOG_LEVEL"); val != "" {
		var levels = map[string]int{
			"debug":    DEBUG,
			"info":     INFO,
			"warning":  WARNING,
			"error":    ERROR,
			"critical": CRITICAL,
			"fatal":    CRITICAL,
		}
		level = levels[val]
	}

	return &loggerImpl{module: module, level: level}
}

// Print a message.
func (l *loggerImpl) print(level int, lstr string, msg string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.level > level {
		return
	}

	t := time.Now().Format("2010-01-02 11:01:01")
	p := fmt.Sprintf("[%s] %s %s ", l.module, t, lstr)
	s := fmt.Sprintf(msg, args...)
	fmt.Println(chalk.Red, p, chalk.Reset, s)
}

// Debug log.
func (l *loggerImpl) Debug(msg string, args ...interface{}) {
	l.print(DEBUG, "DEBU", msg, args...)
}

// Info log.
func (l *loggerImpl) Info(msg string, args ...interface{}) {
	l.print(INFO, "INFO", msg, args...)
}

// Warning log.
func (l *loggerImpl) Warning(msg string, args ...interface{}) {
	l.print(WARNING, "WARN", msg, args...)
}

// Error log.
func (l *loggerImpl) Error(msg string, args ...interface{}) {
	l.print(ERROR, "ERRO", msg, args...)
}

// Fatal log.
func (l *loggerImpl) Fatal(msg string, args ...interface{}) {
	l.print(CRITICAL, "FATA", msg, args...)
}
